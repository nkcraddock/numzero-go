package game

import (
	"errors"
	"sync"
)

// Store stores game shit
type Store interface {
	SavePlayer(p *Player) error
	GetPlayer(email string) (*Player, error)
	ListPlayers() ([]*Player, error)
	SaveRule(r Rule) error
	GetRule(code string) (Rule, error)
}

// NewMemoryStore stands up a super simple in-memory storage
func NewMemoryStore() Store {
	return &memoryStore{
		mu:      new(sync.Mutex),
		players: make(map[string]*Player),
		rules:   make(map[string]Rule),
	}
}

type memoryStore struct {
	mu      *sync.Mutex
	players map[string]*Player
	rules   map[string]Rule
}

func (ms *memoryStore) ListPlayers() ([]*Player, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	players := make([]*Player, len(ms.players))

	i := 0
	for _, p := range ms.players {
		players[i] = p
	}

	return players, nil
}

func (ms *memoryStore) SavePlayer(p *Player) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.players[p.Name] = p
	return nil
}

func (ms *memoryStore) GetPlayer(email string) (*Player, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if p, ok := ms.players[email]; ok {
		return p, nil
	}

	return nil, errors.New("Player not found")
}

func (ms *memoryStore) SaveRule(r Rule) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.rules[r.Code] = r
	return nil
}

func (ms *memoryStore) GetRule(code string) (Rule, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if r, ok := ms.rules[code]; ok {
		return r, nil
	}

	return Rule{}, errors.New("Rule not found")
}
