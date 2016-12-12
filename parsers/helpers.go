package parsers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PositionParser interface {
	Parse(input string) (Position, error)
}
type PlayerParser interface {
	Parse(input string) (Player, error)
}

type Position struct {
	x, y, z int
}

func ParsePosition(input string) (Position, error) {
	var p Position
	parts := strings.Split(input, " ")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return p, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return p, err
	}
	z, err := strconv.Atoi(parts[2])
	if err != nil {
		return p, err
	}
	p = Position{
		x: x,
		y: y,
		z: z,
	}
	return p, nil
}

var playerRegex = regexp.MustCompile(`^(.+)<(\d+)><(.+)><(.*)>$`)

const ErrNoMatch = "Input %v did not match a player"

func ParsePlayer(input string) (Player, error) {
	var p Player
	matches := playerRegex.FindStringSubmatch(input)
	if matches == nil {
		return p, fmt.Errorf(ErrNoMatch, input)
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
