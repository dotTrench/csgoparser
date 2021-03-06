package parsers

import "testing"

func TestParsePurchase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected Props
	}{
		{
			input: `"Kevin<10><BOT><TERRORIST>" purchased "m4a1"`,
			expected: Props{
				"player": Player{
					Name:    "Kevin",
					UID:     10,
					SteamID: "BOT",
					Team:    "TERRORIST",
				},
				"weapon": "m4a1",
			},
		},
	}
	parser := NewPurchasedParser()
	for _, test := range tests {
		p, err := parser.Parse(test.input)
		if err != nil {
			t.Errorf("Got error: %v", err)
		}
		if p["player"] != test.expected["player"] {
			t.Errorf("Player - expected %v got %v", test.expected["player"], p["player"])
		}
		if p["weapon"] != test.expected["weapon"] {
			t.Errorf("Weapon - expected %v got %v", test.expected["weapon"], p["weapon"])
		}
	}
}
