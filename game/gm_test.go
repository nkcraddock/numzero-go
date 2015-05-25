package game_test

import (
	"encoding/json"
	"io/ioutil"

	"github.com/nkcraddock/numzero/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("game.redisStore integration tests", func() {
	var store game.Store
	var gm *game.GM
	var data *testData

	BeforeEach(func() {
		data = loadTestData()
		store = openStore()
		storeTestData(data, store)
		gm = game.NewGameMaster(store)
	})

	AfterEach(func() {
		store.Close()
	})

	Context("AddEvent", func() {
		It("calculates the total score of an event", func() {
			err := gm.AddEvent(data.Events["mervis-1"])
			Ω(err).ShouldNot(HaveOccurred())
			Ω(data.Events["mervis-1"].Total).Should(Equal(4))
		})

		It("persists the event", func() {
			gm.AddEvent(data.Events["mervis-1"])
			evt, err := store.GetEvent(data.Events["mervis-1"].Id)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(evt).Should(Equal(data.Events["mervis-1"]))
		})

		It("returns an error if the player doesnt exist", func() {
			err := gm.AddEvent(data.Events["stubart"])
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(game.ErrorInvalidPlayer))
		})

		It("returns an error if any rules dont exist", func() {
			err := gm.AddEvent(data.Events["chad-badrule"])
			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(game.ErrorInvalidRule))
		})

		It("awards points to the player", func() {
			gm.AddEvent(data.Events["mervis-1"])
			p, err := store.GetPlayer("mervis")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(p.Score).Should(Equal(4))
		})

	})
})

// stores test data loaded from json files
type testData struct {
	Rules   map[string]*game.Rule
	Players map[string]*game.Player
	Events  map[string]*game.Event
}

// loads the test data from the json files
func loadTestData() *testData {
	var players map[string]*game.Player
	data, _ := ioutil.ReadFile("../testdata/test_players.json")
	if err := json.Unmarshal(data, &players); err != nil {
		Ω(err).ShouldNot(HaveOccurred())
		return nil
	}

	var rules map[string]*game.Rule
	data, _ = ioutil.ReadFile("../testdata/test_rules.json")
	if err := json.Unmarshal(data, &rules); err != nil {
		Ω(err).ShouldNot(HaveOccurred())
		return nil
	}

	var events map[string]*game.Event
	data, _ = ioutil.ReadFile("../testdata/test_events.json")
	if err := json.Unmarshal(data, &events); err != nil {
		Ω(err).ShouldNot(HaveOccurred())
		return nil
	}

	return &testData{rules, players, events}
}

// stores the players and rules from the test data into the store
func storeTestData(data *testData, store game.Store) {
	// Add the players from the test data
	for _, p := range data.Players {
		store.SavePlayer(p)
	}
	// Add the rules from the test data
	for _, r := range data.Rules {
		store.SaveRule(r)
	}
}

// sets up a redis store, opens the connection, and wipes the data
func openStore() game.Store {
	rstore, err := game.NewRedisStore("localhost:6379", "", 10)
	Ω(err).ShouldNot(HaveOccurred())
	err = rstore.Open()
	Ω(err).ShouldNot(HaveOccurred())
	rstore.FlushDb()
	return rstore
}
