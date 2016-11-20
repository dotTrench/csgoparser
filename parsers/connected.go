package parsers

import "regexp"

func matchConnected(matches []string) (Props, error) {
	player, err := ParsePlayer(matches[1])
	if err != nil {
		return nil, err
	}
	props := Props{
		"player":  player,
		"address": matches[2],
	}
	return props, nil
}

func NewConnectedParser() LogEventParser {
	return &RegexEventParser{
		regex:   regexp.MustCompile(`^"(.+)"\sconnected,\saddress\s"(.*)"$`),
		matcher: matchConnected,
	}
}
