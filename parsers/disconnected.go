package parsers

import "regexp"

var disconnectedRegex = regexp.MustCompile(`^"(.+)"\sdisconnected\s\(reason\s"(.+)"\)$`)

func ParseDisconnected(input string) (bool, Props, error) {
	var p Props
	matches := disconnectedRegex.FindStringSubmatch(input)

	if matches == nil {
		return false, p, nil
	}

	player, err := ParsePlayer(matches[1])
	if err != nil {
		return true, p, err
	}
	p = Props{
		"player": player,
		"reason": matches[2],
	}

	return true, p, nil
}
