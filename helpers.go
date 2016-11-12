package csgoparser

import (
	"strconv"
	"strings"
)

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
