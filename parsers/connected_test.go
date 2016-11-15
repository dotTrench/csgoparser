package parsers

import "testing"

func TestParseConnected(t *testing.T) {
	tests := []struct {
		input    string
		expected Props
	}{
		{
			input: `"erik<2><STEAM_1:0:7399491><>" connected, address ""`,
			expected: Props{
				"player": Player{
					Name:    "erik",
					UID:     2,
					SteamID: "STEAM_1:0:7399491",
					Team:    "",
				},
				"address": "",
			},
		},
	}
	parser := NewConnectedParser()

	for _, test := range tests {
		got, err := parser.Parse(test.input)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if got["player"] != test.expected["player"] {
			t.Errorf("Player - expected: %v got %v", test.expected["player"], got["player"])
		}
		if got["address"] != test.expected["address"] {
			t.Errorf("Adress - expected: %v got %v", test.expected["address"], got["address"])
		}
	}
}
