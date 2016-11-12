package parsers

type EventParserFunc func(string) (ok bool, props Props, err error)
type EventParser struct {
	EventName string
	Parser    EventParserFunc
}
