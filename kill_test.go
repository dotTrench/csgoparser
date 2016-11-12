package csgoparser

import "testing"

func TestKillParse(t *testing.T) {
	t.Parallel()
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
				"killerPos": Position{
					x: 704,
					y: 2716,
					z: 1,
				},
				"victimPos": Position{
					x: 324,
					y: 3070,
					z: 67,
				},
			},
		}, {
			input: `"Dan<5><BOT><TERRORIST>" [-320 1543 -125] killed "Seth<13><BOT><CT>" [-586 1836 -54] with "ak47"`,
			expected: Props{
				"killer": Player{
					Name:    "Dan",
					UID:     5,
					SteamID: "BOT",
					Team:    "TERRORIST",
				},
				"killerPos": Position{
					x: -320,
					y: 1543,
					z: -125,
				},
				"victim": Player{
					Name:    "Seth",
					UID:     13,
					SteamID: "BOT",
					Team:    "CT",
				},
				"victimPos": Position{
					x: -586,
					y: 1836,
					z: -54,
				},
				"weapon":   "ak47",
				"headshot": false,
			},
		},
	}
	for _, test := range tests {
		ok, props, err := ParseKill(test.input)

		if err != nil {
			t.Errorf("Got error: %v", err)
		}
		if !ok {
			t.Error("Didnt return OK")
		}
		if props == nil {
			t.Error("Props were nil")
		}
		expected := test.expected
		if expected["killer"] != props["killer"] {
			t.Errorf("Killer - expected: %v got %v", expected["killer"], props["killer"])
		}
		if expected["killerPos"] != props["killerPos"] {
			t.Errorf("KillerPos - expected: %v got %v", expected["killerPos"], props["killerPos"])
		}
		if expected["victim"] != props["victim"] {
			t.Errorf("Victim - expected: %v got %v", expected["victim"], props["victim"])
		}
		if expected["victimPos"] != props["victimPos"] {
			t.Errorf("VictimPos - expected: %v got %v", expected["victimPos"], props["victimPos"])
		}
		if expected["weapon"] != props["weapon"] {
			t.Errorf("Weapon - expected: %v got %v", expected["weapon"], props["weapon"])
		}
		if expected["headshot"] != props["headshot"] {
			t.Errorf("Headshot - expected: %v got %v", expected["headshot"], props["headshot"])
		}
	}
}
