package game

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	Image string `json:"image"`
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Image: "http://placehold.it/160x160",
		Score: 0,
	}
}
