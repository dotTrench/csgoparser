package csgoparser

import "regexp"

var purchaseRegex = regexp.MustCompile(`^"(.+)"\spurchased\s"(.+)"$`)

func ParsePurchase(input string) (bool, Props, error) {
	var p Props
	matches := purchaseRegex.FindStringSubmatch(input)

	if matches == nil {
		return false, p, nil
	}
	player, err := ParsePlayer(matches[1])
	if err != nil {
		return true, p, err
	}
	p = Props{
		"player": player,
		"weapon": matches[2],
	}
	return true, p, nil
}
