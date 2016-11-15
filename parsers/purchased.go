package parsers

import "regexp"

func matchPurchase(matches []string) (Props, error) {

	player, err := ParsePlayer(matches[1])
	if err != nil {
		return nil, err
	}
	p := Props{
		"player": player,
		"weapon": matches[2],
	}
	return p, nil
}

func NewPurchasedParser() LogEventParser {
	return &RegexEventParser{
		regex:   regexp.MustCompile(`^"(.+)"\spurchased\s"(.+)"$`),
		matcher: matchPurchase,
	}
}
