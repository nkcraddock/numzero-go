package server

type Activity struct {
	Player      string         `json:"player"`
	Description string         `json:"desc"`
	Url         string         `json:"url"`
	Scores      map[string]int `json:"scores"`
}
