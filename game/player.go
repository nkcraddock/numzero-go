package game

type Rule struct {
	Description string
	Points      int
}

type Player struct {
	Name   string
	Score  int
	Events []Event
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:   name,
		Score:  0,
		Events: make([]Event, 0),
	}
}

func (p *Player) AddEvent(evt Event) error {
	// Store some shit I guess
	p.Score += evt.Score()
	p.Events = append(p.Events, evt)
	return nil
}

type Event struct {
	Rule  Rule
	Count int
}

func (e Event) Score() int {
	return e.Rule.Points * e.Count
}
