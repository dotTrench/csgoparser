package csgoparser

import "time"

type LogEvent struct {
	Timestamp time.Time
	EventType string
	Props     map[string]interface{}
}

const (
	SayEvent  = "say"
	KillEvent = "kill"
)
