package game

// Store stores game shit
type Store interface {
	// Players
	SavePlayer(p *Player) error
	GetPlayer(name string) (*Player, error)
	ListPlayers() ([]*Player, error)

	// Rules
	SaveRule(r *Rule) error
	GetRule(code string) (*Rule, error)
	ListRules() ([]*Rule, error)

	// Events
	SaveEvent(e *Event) error
	GetEvent(id string) (*Event, error)
	GetPlayerEvents(name string, count int64) ([]*Event, error)
}
