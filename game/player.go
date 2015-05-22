package game

type Player struct {
	Name   string  `json:"name"`
	Score  int     `json:"score"`
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
	evt.Total = 0
	for _, s := range evt.Scores {
		evt.Total += s.Rule.Points * s.Times
	}
	p.Score += evt.Total
	p.Events = append(p.Events, *evt)
	return nil
}
