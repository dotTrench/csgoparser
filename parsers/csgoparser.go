package parsers

import (
	"errors"
	"regexp"
	"time"
)

const (
	EventSay          = "say"
	EventKill         = "kill"
	EventDisconnected = "disconnected"
	EventConnected    = "connected"
	EventAssist       = "assist"
	EventPurchased    = "purchased"
)
const ErrTokenizerNoMatch = "Tokenizer no match"

var msgRegx = regexp.MustCompile(`L\s(.+?):\s(.+)$`)

func tokenize(line string) (string, string, error) {
	matches := msgRegx.FindStringSubmatch(line)
	if matches == nil {
		return "", "", errors.New(ErrTokenizerNoMatch)
	}
	return matches[1], matches[2], nil
}

type EventParser struct {
	eventName string
	parser    LogEventParser
}
type CSGOLogParser struct {
	eventParsers []EventParser
}

type LogEvent struct {
	Timestamp  time.Time
	EventType  string
	Properties Props
}

func (p *CSGOLogParser) ParseLine(line string) (LogEvent, error) {
	l := LogEvent{
		EventType:  "unknown",
		Properties: nil,
	}
	timestamp, msg, err := tokenize(line)
	if err != nil {
		return l, err
	}

	tsChan := make(chan time.Time)
	go func(input string) {
		ts, err := ParseTimeStamp(timestamp)
		if err != nil {
			panic(err)
		}
		tsChan <- ts
	}(timestamp)

	for _, e := range p.eventParsers {
		props, err := e.parser.Parse(msg)
		if err != nil {
			return l, err
		}
		if props != nil {
			l.EventType = e.eventName
			l.Properties = props
			break
		}
	}
	l.Timestamp = <-tsChan
	return l, nil
}

func NewParser() *CSGOLogParser {
	return &CSGOLogParser{
		eventParsers: []EventParser{
			{
				eventName: EventSay,
				parser:    NewSayParser(),
			},
			{
				eventName: EventKill,
				parser:    NewKillParser(),
			},
			{
				eventName: EventConnected,
				parser:    NewConnectedParser(),
			},
			{
				eventName: EventAssist,
				parser:    NewAssistParser(),
			},
		},
	}
}
