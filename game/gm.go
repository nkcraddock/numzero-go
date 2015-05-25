package game

type GM struct {
	store        Store
	achievements []Achievement
}

func NewGameMaster(store Store) *GM {
	achievements := []Achievement{
		Achievement{
			Name:        "Too much coffee",
			Description: "drank too much coffee",
			Triggers: map[string]int{
				"coffee": 100,
			},
		},
		Achievement{
			Name:        "Run your mouth",
			Description: "talked a little shit",
			Triggers: map[string]int{
				"shittalking": 25,
			},
		},
		Achievement{
			Name:        "Talkin that shit",
			Description: "talked a whole lot of shit",
			Triggers: map[string]int{
				"shittalking": 100,
			},
		},
		Achievement{
			Name:        "Memey Memerson",
			Description: "https://www.youtube.com/watch?v=ww2e51X_2kw",
			Triggers: map[string]int{
				"memes": 20,
			},
		},
		Achievement{
			Name:        "Oh look who was on the geology dev team",
			Description: "MEEEEEEEEMES",
			Triggers: map[string]int{
				"memes": 50,
			},
		},
	}

	return &GM{store, achievements}
}

// Adds a new event to the game.
func (gm *GM) AddEvent(e *Event) (*EventResult, error) {
	// Get the player
	player, err := gm.store.GetPlayer(e.Player)
	if err != nil {
		return nil, ErrorInvalidPlayer
	}

	// Calculate the total score of the event
	if err := gm.ScoreEvent(e); err != nil {
		return nil, err
	}

	// Credit the player with the event and find out what progress was made
	progress, err := player.AddEvent(e)
	if err != nil {
		return nil, err
	}

	// Check for achievements
	achievements := gm.CheckProgress(progress, player)
	if len(achievements) > 0 {
		// Add any achievements that were triggered
		for _, a := range achievements {
			player.AddAchievement(a)
		}
	}

	// Store the updates to the player
	if err := gm.store.SavePlayer(player); err != nil {
		return nil, err
	}

	// Store the event
	if err := gm.store.SaveEvent(e); err != nil {
		return nil, err
	}

	// Return the results
	return &EventResult{
		Player:       player.Name,
		Points:       e.Total,
		Achievements: achievements,
	}, nil
}

func (gm *GM) CheckProgress(progress AchievementProgress, player *Player) []Achievement {
	var result []Achievement
	for _, a := range gm.achievements {
		if a.CheckProgress(progress, player) {
			result = append(result, a)
		}
	}

	return result
}

// Recalculates the event's total score
func (gm *GM) ScoreEvent(e *Event) error {
	total := 0
	for _, score := range e.Scores {
		rule, err := gm.store.GetRule(score.Rule)
		if err != nil {
			if err == ErrorNotFound {
				return ErrorInvalidRule
			}

			return err
		}

		total += rule.Points * score.Times
	}

	e.Total = total
	return nil
}
