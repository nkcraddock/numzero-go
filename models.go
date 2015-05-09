package numzero

import "time"

type Team struct {
	Name     string             `json:"name"`
	Projects map[string]Project `json:"projects"`
}

type Project struct {
	Name            string `json:"name"`
	BuildConfigName string `json:"build_config_name"`
	Score           int    `json:"score"`
}

type RawBuildData struct {
	Id              string        `json:"id"`
	Date            time.Time     `json:"date"`
	BuildConfigName string        `json:"build_config_name"`
	BuildUrl        string        `json:"build_url"`
	BuildId         string        `json:"build_id"`
	Passed          bool          `json:"passed"`
	Events          RawBuildEvent `json:"events"`
	Participants    []string      `json:"committers"`
}

type RawBuildEvent struct {
	EventType int `json:"event_type"`
	Times     int `json:"times"`
}
