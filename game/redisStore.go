package game

import (
	"encoding/json"
	"strings"

	"gopkg.in/redis.v3"
)

type RedisStore struct {
	opts *redis.Options
	conn *redis.Client
}

const (
	keyPlayers = "players"
	timeout    = 0
)

func NewRedisStore(opts *redis.Options) *RedisStore {
	store := &RedisStore{
		opts: opts,
	}

	return store
}

// SavePlayer will add a new player or update an existing player
func (s *RedisStore) SavePlayer(p *Player) error {
	err := s.connect()
	if err != nil {
		return err
	}
	defer s.close()

	playerJson, err := json.Marshal(p)
	if err != nil {
		return nil
	}

	err = s.conn.HSet(keyPlayers, playerKey(p.Name), string(playerJson)).Err()

	return err
}

// GetPlayer will retrieve an existing player or return ErrorPlayerNotFound
func (s *RedisStore) GetPlayer(name string) (*Player, error) {
	if err := s.connect(); err != nil {
		return nil, err
	}
	defer s.close()

	playerJson, err := s.conn.HGet(keyPlayers, playerKey(name)).Result()

	if err == redis.Nil {
		return nil, ErrorPlayerNotFound
	} else if err != nil {
		return nil, err
	}

	player := new(Player)

	if err = json.Unmarshal([]byte(playerJson), player); err != nil {
		return nil, err
	}

	return player, nil
}

// ListPlayers retrieves a list of all players
func (s *RedisStore) ListPlayers() ([]*Player, error) {
	if err := s.connect(); err != nil {
		return nil, err
	}
	defer s.close()

	results, err := s.conn.HGetAllMap(keyPlayers).Result()

	if err != nil {
		return nil, err
	}

	players := make([]*Player, len(results))

	i := 0
	for _, playerJson := range results {
		p := new(Player)
		if err = json.Unmarshal([]byte(playerJson), p); err != nil {
			return nil, err
		}

		players[i] = p
		i += 1
	}

	return players, nil
}

func (s *RedisStore) connect() error {
	client := redis.NewClient(s.opts)
	if _, err := client.Ping().Result(); err != nil {
		return err
	}

	s.conn = client
	return nil
}

func (s *RedisStore) close() {
	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
	}
}

func playerKey(id string) string {
	id = strings.ToLower(id)
	return id
}

func (s *RedisStore) FlushDb() error {
	if err := s.connect(); err != nil {
		return err
	}

	return s.conn.FlushDb().Err()
}
