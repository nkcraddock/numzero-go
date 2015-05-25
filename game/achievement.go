package game

type Achievement struct {
	// Name is the name (and key) of the achievement.
	Name string `json:"name"`

	// Description is a longer description of the achievement
	Description string `json:"desc"`

	// Trigger is an array of things that will trigger the granting
	// of this achievement to a player
	Triggers map[string]int `json:"triggers"`
}

type AchievementProgress map[string]*RuleProgress

type RuleProgress struct {
	Rule string
	Old  int
	New  int
}

func (a Achievement) CheckProgress(progress AchievementProgress, player *Player) bool {
	// Find out if any of the progress might have crossed a trigger
	relevant := false
	for rule, p := range progress {
		if t, ok := a.Triggers[rule]; ok {
			if p.Old < t && p.New >= t {
				relevant = true
				break
			}
		}
	}

	// if this event didnt cross a trigger, nevermind
	if !relevant {
		return false
	}

	// Find out if ALL triggers have been met by the player
	for rule, t := range a.Triggers {
		p, ok := player.Progress[rule]
		if !ok || p < t {
			return false
		}
	}

	return true
}
