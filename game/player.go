package game

type Player struct {
	Name         string                `json:"name"`
	Score        int                   `json:"score"`
	Image        string                `json:"image"`
	Progress     map[string]int        `json:"progress"`
	Achievements map[string]*Timestamp `json:"achievements"`
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:         name,
		Image:        "http://placehold.it/160x160",
		Score:        0,
		Progress:     make(map[string]int),
		Achievements: make(map[string]*Timestamp),
	}
}

func (p *Player) AddEvent(evt *Event) (AchievementProgress, error) {
	evt.Player = p.Name
	p.Score += evt.Total

	if p.Progress == nil {
		p.Progress = make(map[string]int)
	}

	result := make(AchievementProgress)

	for _, s := range evt.Scores {
		result[s.Rule] = &RuleProgress{Rule: s.Rule}

		if old, ok := p.Progress[s.Rule]; ok {
			result[s.Rule].Old = old
			p.Progress[s.Rule] += s.Times
		} else {
			result[s.Rule].Old = 0
			p.Progress[s.Rule] = s.Times
		}

		result[s.Rule].New = p.Progress[s.Rule]
	}

	return result, nil
}

func (p *Player) AddAchievement(a Achievement) error {
	if p.Achievements == nil {
		p.Achievements = make(map[string]*Timestamp)
	}

	if _, ok := p.Achievements[a.Name]; ok {
		return ErrorDuplicateAchievement
	}

	p.Achievements[a.Name] = NewTimestamp()

	return nil
}
