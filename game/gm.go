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
			Description: "Talking shit about people",
			Triggers: map[string]int{
				"shittalking": 5,
			},
		},
		Achievement{
			Name:        "Talking shit",
			Description: "Talking way too much shit about people",
			Triggers: map[string]int{
				"shittalking": 10,
			},
		},
		Achievement{
			Name:        "Memey Memerson",
			Description: "How original",
			Triggers: map[string]int{
				"memes": 10,
			},
		},
		Achievement{
			Name:        "Oh look who was on the geology dev team",
			Description: "We get it. You know the sayings.",
			Triggers: map[string]int{
				"memes": 50,
			},
		},
		Achievement{
			Name:        "Getting in on the bullshit",
			Description: "It's high time you got in on this bullshit",
			Triggers: map[string]int{
				"memes": 1,
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
