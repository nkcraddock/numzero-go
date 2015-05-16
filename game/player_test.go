package game_test

import (
	"github.com/nkcraddock/numzero/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game tests", func() {
	Context("Player", func() {
		var dude *game.Player
		rule_coffee := game.Rule{"Making a new pot of coffee", 2}
		rule_highfive := game.Rule{"High fiving someone", -1}

		BeforeEach(func() {
			dude = game.NewPlayer("Dude")
		})

		It("gets points for events", func() {
			dude.AddEvent(game.Event{rule_coffee, 2})
			dude.AddEvent(game.Event{rule_highfive, 1})
			Ω(dude.Score).Should(Equal(3))
		})

		It("lists all its events", func() {
			dude.AddEvent(game.Event{rule_coffee, 2})
			dude.AddEvent(game.Event{rule_highfive, 1})
			Ω(dude.Events).Should(HaveLen(2))
		})
	})
})
