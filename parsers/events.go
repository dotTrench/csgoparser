package parsers

import "time"

type Props map[string]interface{}

type LogEvent struct {
	Timestamp time.Time
	EventType string
	Props     Props
}

const (
	SayEvent  = "say"
	KillEvent = "kill"
)
