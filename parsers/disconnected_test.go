package parsers

import "testing"

func TestParseDisconnected(t *testing.T) {
	tests := []struct {
		input    string
		expected Props
	}{
		{
			input: `"Wayne<4><BOT><CT>" disconnected (reason "Kicked by Console")`,
			expected: Props{
				"player": Player{
					Name:    "Wayne",
					UID:     4,
					SteamID: "BOT",
					Team:    "CT",
				},
				"reason": "Kicked by Console",
			},
		},
	}

	for _, test := range tests {
		ok, got, err := ParseDisconnected(test.input)
		if err != nil {
			t.Errorf("Got error: %v", err)
		}
		if !ok {
			t.Error("Didnt get OK")
		}
		want := test.expected
		if got["player"] != want["player"] {
			t.Errorf("Player - expected %v got %v", want["player"], got["player"])
		}
		if got["reason"] != want["reason"] {
			t.Errorf("Reason - expected %v got %v", want["reason"], got["reason"])
		}
	}
}
