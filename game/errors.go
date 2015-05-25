package game

import "errors"

var (
	ErrorNotFound             = errors.New("Not found")
	ErrorInvalidPlayer        = errors.New("Invalid player")
	ErrorInvalidRule          = errors.New("Invalid rule")
	ErrorDuplicateAchievement = errors.New("Player already has that achievement")
)
