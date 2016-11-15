package csgoparser

import (
	"errors"
	"regexp"
	"time"

	"github.com/dotTrench/csgoparser/parsers"
)

const ErrTokenizerNoMatch = "Tokenizer no match"

var msgRegx = regexp.MustCompile(`L\s(.+):\s(.+)$`)

func tokenize(line string) (string, string, error) {
	matches := msgRegx.FindStringSubmatch(line)
	if matches == nil {
		return "", "", errors.New(ErrTokenizerNoMatch)
	}
	return matches[1], matches[2], nil
}

type CSGOLogParser struct {
	eventParsers []parsers.LogEventParser
}

type LogEvent struct {
	Timestamp  time.Time
	EventType  string
	Properties parsers.Props
}

func (p *CSGOLogParser) ParseLine(line string) (LogEvent, error) {
	var l LogEvent
	timestamp, msg, err := tokenize(line)
	if err != nil {
		return l, err
	}
	ts, err := parsers.ParseTimeStamp(timestamp)
	if err != nil {
		return l, err
	}
	l.Timestamp = ts
	for _, e := range p.eventParsers {
		props, err := e.Parse(msg)
		if err != nil {
			return l, err
		}
		if props == nil {
			continue
		}
		l.Properties = props
	}
	return l, nil
}
