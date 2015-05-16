package numzero

import "github.com/dgrijalva/jwt-go"

type Store interface {
	SaveSession(string, *jwt.Token)
	GetSession(string) (*jwt.Token, bool)
}
