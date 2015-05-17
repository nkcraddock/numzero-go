package game

type Event struct {
	Rule  Rule
	Count int
}

func (e Event) Score() int {
	return e.Rule.Points * e.Count
}
