package game_test

import (
	"github.com/nkcraddock/numzero/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/redis.v3"
)

var _ = Describe("game.redisStore integration tests", func() {
	var store *game.RedisStore

	BeforeEach(func() {
		options := &redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       10,
		}

		store = game.NewRedisStore(options)
		store.FlushDb()
	})

	Context("Players", func() {
		chad := &game.Player{Name: "Chad"}

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
			_, err := store.GetPlayer("totally made up")
			Ω(err).Should(Equal(game.ErrorPlayerNotFound))
		})

	})
})
