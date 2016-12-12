package receiver

type Receiver interface {
	GetLine() (string, error)
}
