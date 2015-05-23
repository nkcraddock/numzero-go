package game

// Store stores game shit
type Store interface {
	SavePlayer(p *Player) error
	GetPlayer(name string) (*Player, error)
	ListPlayers() ([]*Player, error)
	SaveRule(r *Rule) error
	GetRule(code string) (*Rule, error)
	ListRules() ([]*Rule, error)
}
