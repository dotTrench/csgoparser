package parsers

import "regexp"

var sayRegex = regexp.MustCompile(`"(.+)"\ssay\s(.+)$`)

func ParseSay(input string) (bool, Props, error) {
	var p Props
	matches := sayRegex.FindStringSubmatch(input)
	if matches == nil {
		return false, p, nil
	}
	player, err := ParsePlayer(matches[1])

	if err != nil {
		return false, p, err
	}

	p = Props{
		"sender":  player,
		"message": matches[2],
	}
	return true, p, nil
}
