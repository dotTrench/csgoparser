package csgoparser

import "testing"

func TestParsePosition(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected Position
	}{
		{
			input: "704 2716 1",
			expected: Position{
				x: 704,
				y: 2716,
				z: 1,
			},
		},
		{
			input: "-132 193 -1",
			expected: Position{
				x: -132,
				y: 193,
				z: -1,
			},
		},
	}

	for _, test := range tests {
		p, err := ParsePosition(test.input)

		if err != nil {
			t.Errorf("Got error: %v", err)
		}
		if p != test.expected {
			t.Errorf("Expected %v got %v", test.expected, p)
		}
	}
}
