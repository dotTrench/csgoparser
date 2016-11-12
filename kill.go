package csgoparser

import "regexp"

var killRegex = regexp.MustCompile(``)

func ParseKill(input string) (bool, Props, error) {
	var p Props
	return false, p, nil
}
