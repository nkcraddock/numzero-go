package game

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Score: 0,
	}
}

func (p *Player) AddEvent(evt *Event) error {
	evt.Player = p.Name
	p.Score += evt.Total
	return nil
}
