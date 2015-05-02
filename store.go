package gooby

type Store struct {
	companies map[string]Company
}

func NewStore(companies ...string) *Store {
	s := &Store{companies: make(map[string]Company)}
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
