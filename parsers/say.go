package parsers

import "regexp"

func matchSay(matches []string) (Props, error) {
	player, err := ParsePlayer(matches[1])

	if err != nil {
		return nil, err
	}

	p := Props{
		"sender":  player,
		"message": matches[2],
	}

	return p, nil
}

func NewSayParser() LogEventParser {
	return &RegexEventParser{
		regex:   regexp.MustCompile(`^"(.+)"\ssay\s"(.+)"$`),
		matcher: matchSay,
	}
}
