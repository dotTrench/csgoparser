package parsers

import "regexp"

var assistRegex = regexp.MustCompile(`^"(.+)"\sassisted\skilling\s"(.+)"$`)

func matchAssist(matches []string) (Props, error) {

	assister, err := ParsePlayer(matches[1])
	if err != nil {
		return nil, err
	}
	victim, err := ParsePlayer(matches[2])
	if err != nil {
		return nil, err
	}
	p := Props{
		"assister": assister,
		"victim":   victim,
	}

	return p, nil
}
func NewAssistParser() LogEventParser {
	return &RegexEventParser{
		regex:   regexp.MustCompile(`^"(.+)"\sassisted\skilling\s"(.+)"$`),
		matcher: matchAssist,
	}
}
