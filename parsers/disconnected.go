package parsers

import "regexp"

func NewDisconnectedParser() LogEventParser {
	return &RegexEventParser{
		regex:   regexp.MustCompile(`^"(.+)"\sdisconnected\s\(reason\s"(.+)"\)$`),
		matcher: matchDisconnected,
	}
}

func matchDisconnected(matches []string) (Props, error) {
	player, err := ParsePlayer(matches[1])
	if err != nil {
		return nil, err
	}
	p := Props{
		"player": player,
		"reason": matches[2],
	}

	return p, nil
}
