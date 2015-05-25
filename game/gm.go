package game

type GM struct {
	store Store
}

func NewGameMaster(store Store) *GM {
	return &GM{store}
}

func (gm *GM) AddEvent(e *Event) error {
	if _, err := gm.store.GetPlayer(e.Player); err != nil {
		if err == ErrorNotFound {
			return ErrorInvalidPlayer
		}

		return err
	}

	if err := gm.ScoreEvent(e); err != nil {
		return err
	}

	return gm.store.SaveEvent(e)
}

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
