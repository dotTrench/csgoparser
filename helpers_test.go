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

func TestParsePlayer(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  Player
	}{
		{"Console<0><Console><Console>", Player{
			Name:    "Console",
			UID:     0,
			SteamID: "Console",
			Team:    "Console",
		}},
		{"Erik<13><STEAM_0_1:123><TERRORIST>", Player{
			Name:    "Erik",
			UID:     13,
			SteamID: "STEAM_0_1:123",
			Team:    "TERRORIST",
		}},
		{"erik<2><STEAM_1:0:7399491><>", Player{
			Name:    "erik",
			UID:     2,
			SteamID: "STEAM_1:0:7399491",
			Team:    "",
		}},
	}

	for _, test := range tests {
		got, err := ParsePlayer(test.input)
		want := test.want
		if err != nil {
			t.Errorf("Got error: %v", err)
			break
		}
		if got.Name != want.Name {
			t.Errorf("Name - expected %v got %v", want.Name, got.Name)
		}
		if got.SteamID != want.SteamID {
			t.Errorf("SteamID - expected %v got %v", want.SteamID, got.SteamID)
		}
		if got.Team != want.Team {
			t.Errorf("Team - expected %v got %v", want.Team, got.Team)
		}
		if got.UID != want.UID {
			t.Errorf("UID - expected %v got %v", want.UID, got.UID)
		}
	}
}
