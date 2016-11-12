package csgoparser

import "regexp"

var assistRegex = regexp.MustCompile(`^"(.+)"\sassisted\skilling\s"(.+)"$`)

func ParseAssist(input string) (bool, Props, error) {
	var p Props
	matches := assistRegex.FindStringSubmatch(input)
	if matches == nil {
		return false, p, nil
	}
	assister, err := ParsePlayer(matches[1])
	if err != nil {
		return true, p, err
	}
	victim, err := ParsePlayer(matches[2])
	if err != nil {
		return true, p, err
	}
	p = Props{
		"assister": assister,
		"victim":   victim,
	}

	return true, p, nil
}
