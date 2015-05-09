package numzero

import "github.com/dgrijalva/jwt-go"

type MemoryStore struct {
	teams    map[string]Team
	sessions map[string]jwt.Token
}

func NewMemoryStore(teams ...string) Store {
	s := &MemoryStore{
		teams:    make(map[string]Team),
		sessions: make(map[string]jwt.Token),
	}
	for _, c := range teams {
		s.teams[c] = Team{Name: c}
	}
	return s
}

func (s *MemoryStore) GetTeams() []Team {
	teams := make([]Team, len(s.teams))
	i := 0
	for _, c := range s.teams {
		teams[i] = c
		i += 1
	}

	return teams
}

func (s *MemoryStore) SaveTeam(c *Team) {
	s.teams[c.Name] = *c
}

func (s *MemoryStore) DeleteTeam(name string) bool {
	if _, ok := s.teams[name]; !ok {
		return false
	} else {
		delete(s.teams, name)
		return true
	}
}

func (s *MemoryStore) GetTeam(name string) (c Team, ok bool) {
	c, ok = s.teams[name]
	return
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
