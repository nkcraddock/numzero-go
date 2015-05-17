package game

// An Event represents something that has happened
// resulting in Scores
type Event struct {
	// The Description is a short description of the event that has happened
	// ex. "Static analysis on commit 34fe6a... to numzero/master"
	Description string

	// The Url is an optional url to a resource with more details about this
	// Event. ex. "https://github.com/nkcraddock/numzero/commits/81570f37"
	Url string

	// Scores are the matched rules that occurred as part of this event
	// such as Build Failure, Tests Added, etc.
	Scores []Score

	// Total is the calculated total net value in points that will be awarded
	// as a result of this event. It is the total value of all the Scores
	Total int
}

// A Score is a scoring action in an Event where a game Rule was matched
// one or more times. An example would be adding 5 new Unit tests.
// score := &Score{Rule: rule_unittests, 5}
type Score struct {
	// The matching Rule that this Score resulted from
	Rule *Rule

	// The number of times the Rule was triggered in the Event
	Times int
}
