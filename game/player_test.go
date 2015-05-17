package game_test

import (
	"github.com/nkcraddock/numzero/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player", func() {
	var dude *game.Player
	rule_coffee := game.Rule{"coffee", "Making a new pot of coffee", 2}
	rule_highfive := game.Rule{"highfive", "High fiving someone", -1}

	BeforeEach(func() {
		dude = game.NewPlayer("Dude")
	})

	Context("AddEvent", func() {
		It("adds points for events", func() {
			dude.AddEvent(game.Event{rule_coffee, 2})
			dude.AddEvent(game.Event{rule_highfive, 1})
			Ω(dude.Score).Should(Equal(3))
		})

		It("maintains a list of added events", func() {
			dude.AddEvent(game.Event{rule_coffee, 2})
			dude.AddEvent(game.Event{rule_highfive, 1})
			Ω(dude.Events).Should(HaveLen(2))
		})
	})

	Context("Persistence", func() {
		var store game.Store

		BeforeEach(func() {
			store = game.NewMemoryStore()
		})

		It("persists a player", func() {
			dude.AddEvent(game.Event{rule_coffee, 2})
			dude.AddEvent(game.Event{rule_highfive, 1})

			err := store.SavePlayer(dude)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("retrieves a persisted player", func() {
			dude.AddEvent(game.Event{rule_coffee, 2})
			dude.AddEvent(game.Event{rule_highfive, 1})

			err := store.SavePlayer(dude)
			Ω(err).ShouldNot(HaveOccurred())

			newDude, err := store.GetPlayer(dude.Name)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(newDude.Score).Should(Equal(dude.Score))
		})

		It("persists a rule", func() {
			err := store.SaveRule(rule_coffee)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("retrieves a persisted rule", func() {
			err := store.SaveRule(rule_coffee)
			Ω(err).ShouldNot(HaveOccurred())

			newRule, err := store.GetRule(rule_coffee.Code)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(newRule.Description).Should(Equal(rule_coffee.Description))
		})
	})
})