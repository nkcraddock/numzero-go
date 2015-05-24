package game_test

import (
	"time"

	"github.com/nkcraddock/numzero/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("game.redisStore integration tests", func() {
	var store *game.RedisStore
	var evt_won *game.Event
	chad := &game.Player{Name: "Chad"}
	roger := &game.Player{Name: "Roger"}

	got_powerup := &game.Rule{"powerup", "got the powerup", 5}
	won_thegame := &game.Rule{"wonthegame", "won the game", 20}

	BeforeEach(func() {
		var err error
		store, err = game.NewRedisStore("localhost:6379", "", 10)
		Ω(err).ShouldNot(HaveOccurred())

		err = store.Open()
		Ω(err).ShouldNot(HaveOccurred())

		store.FlushDb()

		evt_won = &game.Event{
			Player:      "chad",
			Description: "played the game",
			Scores: []game.Score{
				game.Score{"powerup", 5},
				game.Score{"wonthegame", 6},
			},
			Date: time.Now(),
		}
	})

	AfterEach(func() {
		store.Close()
	})

	Context("Events", func() {
		It("saves an event", func() {
			err := store.SaveEvent(evt_won)
			Ω(err).ShouldNot(HaveOccurred())
		})
		It("retrieves an event", func() {
			store.SaveEvent(evt_won)

			evt, err := store.GetEvent(evt_won.Id)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(evt).ShouldNot(BeNil())
			Ω(evt.Scores).Should(HaveLen(2))
		})
		It("adds the event id to the player's events", func() {
			store.SaveEvent(evt_won)

			evts, err := store.GetPlayerEvents("chad", 0)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(evts).Should(HaveLen(1))
			Ω(evts[0]).Should(Equal(evt_won))
		})
	})

	Context("Rules", func() {
		It("saves a rule", func() {
			err := store.SaveRule(got_powerup)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("retrieves a rule", func() {
			store.SaveRule(got_powerup)

			r, err := store.GetRule("powerup")
			Ω(err).ShouldNot(HaveOccurred())

			Ω(r).ShouldNot(BeNil())
			Ω(r.Points).Should(Equal(5))
		})

		It("retrieves a list of rules", func() {
			store.SaveRule(got_powerup)
			store.SaveRule(won_thegame)

			r, err := store.ListRules()
			Ω(err).ShouldNot(HaveOccurred())

			Ω(r).ShouldNot(BeNil())
			Ω(r).Should(HaveLen(2))
		})
	})

	Context("Players", func() {
		It("saves a player", func() {
			err := store.SavePlayer(chad)
			Ω(err).ShouldNot(HaveOccurred())

			p, err := store.GetPlayer("chad")
			Ω(err).ShouldNot(HaveOccurred())

			Ω(p).ShouldNot(BeNil())
			Ω(p.Name).Should(Equal("Chad"))
		})

		It("retrieves a player", func() {
			store.SavePlayer(chad)

			p, err := store.GetPlayer("chad")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(p).ShouldNot(BeNil())
		})

		It("returns an error if player doesnt exist", func() {
			_, err := store.GetPlayer("chad")
			Ω(err).Should(Equal(game.ErrorNotFound))
		})

		It("retrieves a list of players", func() {
			store.SavePlayer(chad)
			store.SavePlayer(roger)

			players, err := store.ListPlayers()
			Ω(err).ShouldNot(HaveOccurred())
			Ω(players).Should(HaveLen(2))
		})

	})
})
