package gooby

import "github.com/dgrijalva/jwt-go"

type Store interface {
	GetTeams() []Team
	SaveTeam(c *Team)
	DeleteTeam(name string) bool
	GetTeam(name string) (Team, bool)
	SaveSession(string, *jwt.Token)
	GetSession(string) (*jwt.Token, bool)
}

type Repo interface {
	List() []interface{}
	Get(string) (interface{}, error)
	Delete(string) error
	Save(interface{}) error
}
