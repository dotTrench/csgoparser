package csgoparser

import "testing"

func TestParsePlayer(t *testing.T) {
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
