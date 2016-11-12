package csgoparser

import (
	"testing"
	"time"
)

func TestTimestampParser(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{"11/10/2015 - 17:39:44", time.Date(2015, time.November, 10, 17, 39, 44, 0, time.UTC)},
		{"11/10/2015 - 18:59:11", time.Date(2015, time.November, 10, 18, 59, 11, 0, time.UTC)},
	}

	for _, test := range tests {
		got, err := ParseTimeStamp(test.input)
		if err != nil {
			t.Errorf("Got error: %v", err)
			break
		}
		if got != test.want {
			t.Errorf("Expected: %v got %v", test.want, got)
		}
	}
}
