package parsers

import "regexp"

func matchKill(matches []string) (Props, error) {
	killer, err := ParsePlayer(matches[1])
	if err != nil {
		return nil, err
	}
	killerPos, err := ParsePosition(matches[2])
	if err != nil {
		return nil, err
	}
	victim, err := ParsePlayer(matches[3])
	if err != nil {
		return nil, err
	}
	victimPos, err := ParsePosition(matches[4])
	if err != nil {
		return nil, err
	}

	weapon := matches[5]
	headshot := matches[6] == "(headshot)"

	p := Props{
		"killer":    killer,
		"killerPos": killerPos,
		"victim":    victim,
		"victimPos": victimPos,
		"weapon":    weapon,
		"headshot":  headshot,
	}
	return p, nil
}

func NewKillParser() *RegexEventParser {
	return &RegexEventParser{
		regex:   regexp.MustCompile(`"(.+)"\s\[(.+)\]\skilled\s"(.+)"\s\[(.+)\]\swith\s"(.+)"\s?(\(headshot\))?`),
		matcher: matchKill,
	}
}
