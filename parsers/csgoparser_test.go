package parsers

import "testing"

func TestTokenizer(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string

		timestamp string
		message   string
	}{
		{
			`L 11/10/2015 - 18:59:00: "Gabe<7><BOT><CT>" disconnected (reason "Kicked by Console")`,
			`11/10/2015 - 18:59:00`,
			`"Gabe<7><BOT><CT>" disconnected (reason "Kicked by Console")`,
		},
		{
			`L 11/10/2015 - 18:58:37: "Adrian<3><BOT><CT>" [1264 2013 10] killed "Hank<11><BOT><TERRORIST>" [692 2328 68] with "m4a1"`,
			`11/10/2015 - 18:58:37`,
			`"Adrian<3><BOT><CT>" [1264 2013 10] killed "Hank<11><BOT><TERRORIST>" [692 2328 68] with "m4a1"`,
		},
		{
			`L 09/05/2015 - 19:13:45: rcon from "127.0.0.1:55679": command "status"`,
			`09/05/2015 - 19:13:45`,
			`rcon from "127.0.0.1:55679": command "status"`,
		},
	}

	for _, test := range tests {
		ts, msg, err := tokenize(test.input)
		if err != nil {
			t.Errorf("Got error: %v", err)
		}
		if msg != test.message {
			t.Errorf("Expected message: %v got %v", test.message, msg)
		}
		if ts != test.timestamp {
			t.Errorf("Expected timestamp: %v got %v", test.timestamp, ts)
		}
	}
}
