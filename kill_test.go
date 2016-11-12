package csgoparser

import "testing"

func TestKillParse(t *testing.T) {
	tests := []struct {
		input    string
		expected Props
	}{
		{
			input: `L 11/10/2015 - 17:39:49: "erik<2><STEAM_1:0:7399491><TERRORIST>" [704 2716 1] killed "Tyler<9><BOT><CT>" [324 3070 67] with "ak47" (headshot)`,
			expected: Props{
				"killer": Player{
					Name:    "erik",
					UID:     2,
					SteamID: "STEAM_1:0:7399491",
					Team:    "TERRORIST",
				},
				"victim": Player{
					Name:    "Tyler",
					UID:     9,
					SteamID: "BOT",
					Team:    "CT",
				},
				"weapon":   "ak47",
				"headshot": true,
				"killerpos": Position{
					x: 704,
					y: 2716,
					z: 1,
				},
			},
		},
	}
	for i, _ := range tests {
		t.Log(i)
	}
}
