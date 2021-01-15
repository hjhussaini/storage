package dbx

import "github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"

func Root(token string) *dropbox.Config {
	return &dropbox.Config{
		token,
		dropbox.LogOff,
		nil,
		"",
		"",
		nil,
		nil,
		nil,
	}
}
