package numzero

import "github.com/dgrijalva/jwt-go"

type MemoryStore struct {
	sessions map[string]jwt.Token
}

func NewMemoryStore() Store {
	s := &MemoryStore{
		sessions: make(map[string]jwt.Token),
	}
	return s
}

func (s *MemoryStore) GetSession(id string) (*jwt.Token, bool) {
	if token, ok := s.sessions[id]; ok {
		return &token, ok
	}

	return nil, false
}

func (s *MemoryStore) SaveSession(id string, token *jwt.Token) {
	s.sessions[id] = *token
}
