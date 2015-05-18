package game

type Player struct {
	Name   string
	Score  int
	Events []Event `json:"-"`
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:   name,
		Score:  0,
		Events: make([]Event, 0),
	}
}

func (p *Player) AddEvent(evt *Event) error {
	p.Score += evt.Total
	p.Events = append(p.Events, *evt)
	return nil
}
