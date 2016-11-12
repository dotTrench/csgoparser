package parsers

import "testing"

func TestParseAssist(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected Props
	}{
		{
			input: `"Troy<12><BOT><CT>" assisted killing "Ted<8><BOT><TERRORIST>"`,
			expected: Props{
				"assister": Player{
					Name:    "Troy",
					UID:     12,
					SteamID: "BOT",
					Team:    "CT",
				},
				"victim": Player{
					Name:    "Ted",
					UID:     8,
					SteamID: "BOT",
					Team:    "TERRORIST",
				},
			},
		},
	}

	for _, test := range tests {
		ok, got, err := ParseAssist(test.input)

		if err != nil {
			t.Errorf("Got error: %v", err)
		}
		if !ok {
			t.Error("Did not return OK")
		}
		want := test.expected
		if got["killer"] != want["killer"] {
			t.Errorf("Killer - expected: %v got %v", want["killer"], got["killer"])
		}
		if got["victim"] != want["victim"] {
			t.Errorf("Victim - expected: %v got %v", want["victim"], got["victim"])
		}
	}
}
