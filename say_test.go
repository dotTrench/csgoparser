package csgoparser

import "testing"

func TestParseSay(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected Props
	}{
		{
			input: `"Console<0><Console><Console>" say "hello"`,
			expected: Props{
				"sender": Player{
					Name:    "Console",
					UID:     0,
					SteamID: "Console",
					Team:    "Console",
				},
				"message": "hello",
			},
		}, {
			input: `"erik<2><STEAM_1:0:7399491><CT>" say testing "this""`,
			expected: Props{
				"sender": Player{
					Name:    "erik",
					UID:     2,
					SteamID: "STEAM_1:0:7399491",
					Team:    "CT",
				},
				"message": `testing "this"`,
			},
		}, {
			input: `"Console<0><Console><Console>" say ""lol""`,
			expected: Props{
				"sender": Player{
					Name:    "Console",
					UID:     0,
					SteamID: "Console",
					Team:    "Console",
				},
				"message": `"lol"`,
			},
		},
	}

	for _, test := range tests {
		ok, props, err := ParseSay(test.input)

		if !ok {
			t.Errorf("Did not return OK")
		}
		if props == nil {
			t.Errorf("Props were nil")
		}
		if err != nil {
			t.Errorf("Got error: %v", err)
		}

		if props["sender"] != test.expected["sender"] {
			t.Errorf("Sender - expected: %v got %v", test.expected["sender"], props["sender"])
		}
	}
}
