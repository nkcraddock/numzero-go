package gooby

import "github.com/dgrijalva/jwt-go"

type Store struct {
	companies map[string]Company
	sessions  map[string]jwt.Token
}

func NewStore(companies ...string) *Store {
	s := &Store{
		companies: make(map[string]Company),
		sessions:  make(map[string]jwt.Token),
	}
	for _, c := range companies {
		s.companies[c] = Company{Name: c}
	}
	return s
}

func (s *Store) GetCompanies() []Company {
	companies := make([]Company, len(s.companies))
	i := 0
	for _, c := range s.companies {
		companies[i] = c
		i += 1
	}

	return companies
}

func (s *Store) SaveCompany(c *Company) {
	s.companies[c.Name] = *c
}

func (s *Store) DeleteCompany(name string) bool {
	if _, ok := s.companies[name]; !ok {
		return false
	} else {
		delete(s.companies, name)
		return true
	}
}

func (s *Store) GetCompany(name string) (c Company, ok bool) {
	c, ok = s.companies[name]
	return
}

func (s *Store) GetSession(id string) (*jwt.Token, bool) {
	if token, ok := s.sessions[id]; ok {
		return &token, ok
	}

	return nil, false
}

func (s *Store) SaveSession(id string, token *jwt.Token) {
	s.sessions[id] = *token
}
