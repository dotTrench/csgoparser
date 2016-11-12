package csgoparser

import "regexp"

var killRegex = regexp.MustCompile(`"(.+)"\s\[(.+)\]\skilled\s"(.+)"\s\[(.+)\]\swith\s"(.+)"\s?(\(headshot\))?`)

func ParseKill(input string) (bool, Props, error) {
	var p Props
	matches := killRegex.FindStringSubmatch(input)
	if matches == nil {
		return false, p, nil
	}
	killer, err := ParsePlayer(matches[1])
	if err != nil {
		return true, p, err
	}
	killerPos, err := ParsePosition(matches[2])
	if err != nil {
		return true, p, err
	}
	victim, err := ParsePlayer(matches[3])
	if err != nil {
		return true, p, err
	}
	victimPos, err := ParsePosition(matches[4])
	if err != nil {
		return true, p, err
	}

	weapon := matches[5]
	headshot := matches[6] == "(headshot)"

	p = Props{
		"killer":    killer,
		"killerPos": killerPos,
		"victim":    victim,
		"victimPos": victimPos,
		"weapon":    weapon,
		"headshot":  headshot,
	}
	return true, p, nil
}
