package requests

import "time"

type Requests struct {
	Tags          string `json:"tags"` //format: foo=bar,baz=quux etc
	Duration      string `json:"duration"`
	Duration_time time.Duration
	Rate          int     `json:"rate"`
	PriorityBoost int32   `json:"priority_boost"`
	Weight        float32 `json:"weight"`
	ParsedTags    map[string]string
	Started       *time.Time
	Ended         *time.Time
	APIkey        string `json:"apikey"`
	Environment   string `json:"environment"`
	Guard         string `json:"guard"`
}
