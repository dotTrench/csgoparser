package csgoparser

import "time"

const timestampFormat = "01/02/2006 - 15:04:05"

type TimestampParser func(string) (Timestamp, error)
type Timestamp time.Time

func ParseTimeStamp(input string) (time.Time, error) {
	return time.Parse(timestampFormat, input)
}
