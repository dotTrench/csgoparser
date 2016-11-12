package csgoparser

import (
	"errors"
	"regexp"
	"strconv"
)

type Player struct {
	UID     int
	Name    string
	SteamID string
	Team    string
}

var playerRegex = regexp.MustCompile(`^(.+)<(\d+)><(.+)><(.*)>$`)

const ErrNoMatch = "No match"

func ParsePlayer(input string) (Player, error) {
	var p Player
	matches := playerRegex.FindStringSubmatch(input)
	if matches == nil {
		return p, errors.New(ErrNoMatch)
	}

	uid, err := strconv.Atoi(matches[2])
	if err != nil {
		return p, err
	}
	p = Player{
		Name:    matches[1],
		UID:     uid,
		SteamID: matches[3],
		Team:    matches[4],
	}
	return p, nil
}
