package game

import (
	"encoding/json"
	"fmt"
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

	err = s.conn.Set(playerKey(p.Name), string(playerJson), timeout).Err()

	return err
}

// GetPlayer will retrieve an existing player or return ErrorPlayerNotFound
func (s *RedisStore) GetPlayer(name string) (*Player, error) {
	if err := s.connect(); err != nil {
		return nil, err
	}
	defer s.close()

	playerJson, err := s.conn.Get(playerKey(name)).Result()

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
	return fmt.Sprintf("player:%s", id)
}

func (s *RedisStore) FlushDb() error {
	if err := s.connect(); err != nil {
		return err
	}

	return s.conn.FlushDb().Err()
}
