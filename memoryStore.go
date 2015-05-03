package gooby

import "github.com/dgrijalva/jwt-go"

type MemoryStore struct {
	companies map[string]Company
	sessions  map[string]jwt.Token
}

func NewMemoryStore(companies ...string) Store {
	s := &MemoryStore{
		companies: make(map[string]Company),
		sessions:  make(map[string]jwt.Token),
	}
	for _, c := range companies {
		s.companies[c] = Company{Name: c}
	}
	return s
}

func (s *MemoryStore) GetCompanies() []Company {
	companies := make([]Company, len(s.companies))
	i := 0
	for _, c := range s.companies {
		companies[i] = c
		i += 1
	}

	return companies
}

func (s *MemoryStore) SaveCompany(c *Company) {
	s.companies[c.Name] = *c
}

func (s *MemoryStore) DeleteCompany(name string) bool {
	if _, ok := s.companies[name]; !ok {
		return false
	} else {
		delete(s.companies, name)
		return true
	}
}

func (s *MemoryStore) GetCompany(name string) (c Company, ok bool) {
	c, ok = s.companies[name]
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
