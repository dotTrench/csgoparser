package parsers

import (
	"regexp"
)

type Props map[string]interface{}

type RegexpMatcherFunc func([]string) (Props, error)
type RegexEventParser struct {
	regex   *regexp.Regexp
	matcher RegexpMatcherFunc
}

// LogEventParser is responsible for parsing the message part of the event
type LogEventParser interface {
	// input is the raw message string, returns nil props if no match
	Parse(input string) (Props, error)
}

func (parser *RegexEventParser) Parse(input string) (Props, error) {
	matches := parser.regex.FindStringSubmatch(input)
	if matches == nil {
		return nil, nil
	}

	return parser.matcher(matches)
}
