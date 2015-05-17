package server

type Activity struct {
	Description string         `json:"desc"`
	Url         string         `json:"url"`
	Scores      map[string]int `json:"scores"`
}
