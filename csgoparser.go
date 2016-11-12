package csgoparser

import (
	"errors"
	"regexp"
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

type PropExtractorFunc func(string) map[string]interface{}
type PropertyExtractor struct {
	eventName string
	extractor PropExtractorFunc
}
type CSGOLogParser struct {
	extractors []PropertyExtractor
}

func NewParser() CSGOLogParser {
	return CSGOLogParser{
		extractors: []PropertyExtractor{
			{
				eventName: "say",
				extractor: SayParse,
			},
		},
	}
}

func SayParse(line string) map[string]interface{} {
	var m = make(map[string]interface{})
	m["cock"] = 1
	m["sender"] = Player{
		UID:     13,
		Name:    "lolöööö",
		SteamID: "STEAM:01203213",
	}
	return m
}
func (p *CSGOLogParser) ParseLine(line string) (LogEvent, error) {
	var l LogEvent
	timestamp, msg, err := tokenize(line)
	if err != nil {
		return l, err
	}
	ts, err := ParseTimeStamp(timestamp)
	if err != nil {
		return l, err
	}
	l.Timestamp = ts
	for _, e := range p.extractors {
		l.EventType = e.eventName
		l.Props = e.extractor(msg)
	}
	return l, nil
}
