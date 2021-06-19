package main

import "testing"

func TestExtractPath(t *testing.T) {
	tests := []struct {
		URL  string
		Want string
	}{
		{
			URL:  "https://github.com/awesome-streamers/awesome-streamerrc.git",
			Want: "github.com/awesome-streamers/awesome-streamerrc",
		},
		{
			URL:  "https://github.com/awesome-streamers/awesome-streamerrc",
			Want: "github.com/awesome-streamers/awesome-streamerrc",
		},
		{
			URL:  "git@github.com:awesome-streamers/awesome-streamerrc.git",
			Want: "github.com/awesome-streamers/awesome-streamerrc",
		},
	}
	for _, test := range tests {
		got, err := ExtractPath(test.URL)
		if err != nil {
			t.Fatal(err)
		} else if got != test.Want {
			t.Errorf("got=%v want=%v", got, test.Want)
		}
	}
}
