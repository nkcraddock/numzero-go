package game_test

import (
	"github.com/nkcraddock/numzero/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player", func() {
	var dude *game.Player
	var event_one *game.Event
	BeforeEach(func() {
		dude = game.NewPlayer("Dude")

		event_one = &game.Event{
			Description: "Daily Summary",
			Url:         "",
			Total:       3,
			Scores: []game.Score{
				game.Score{"coffee", 2},
				game.Score{"highfive", 1},
			},
		}
	})

	Context("AddEvent", func() {
		It("adds points for events", func() {
			dude.AddEvent(event_one)
			Ω(dude.Score).Should(Equal(3))
		})

		It("sets this player as the event's player", func() {
			dude.AddEvent(event_one)
			Ω(event_one.Player).Should(Equal(dude.Name))
		})
	})
})
