package main

import (
	"fmt"

	"github.com/dotTrench/csgoparser/parsers"
)

type Receiver interface {
	GetLine() (string, error)
}
type Parser interface {
	ParseLine(string) (parsers.LogEvent, error)
}
type Dispatcher interface {
	Dispatch(parsers.LogEvent) error
}

func main() {
	parser := parsers.NewParser()

	event, err := parser.ParseLine("L 2014)")
	if err != nil {
		panic(err)
	}
	fmt.Println(event)
}
