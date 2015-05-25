package game

type GM struct {
	store Store
}

func NewGameMaster(store Store) *GM {
	return &GM{store}
}

// Adds a new event to the game.
func (gm *GM) AddEvent(e *Event) error {
	player, err := gm.store.GetPlayer(e.Player)
	if err != nil {
		if err == ErrorNotFound {
			return ErrorInvalidPlayer
		}

		return err
	}

	if err := gm.ScoreEvent(e); err != nil {
		return err
	}

	player.Score += e.Total
	gm.store.SavePlayer(player)

	return gm.store.SaveEvent(e)
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
