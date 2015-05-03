package gooby

import "github.com/dgrijalva/jwt-go"

type Store interface {
	GetCompanies() []Company
	SaveCompany(c *Company)
	DeleteCompany(name string) bool
	GetCompany(name string) (Company, bool)
	SaveSession(string, *jwt.Token)
	GetSession(string) (*jwt.Token, bool)
}
