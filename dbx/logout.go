package dbx

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/auth"
)

func Logout(configuration dropbox.Config) error {
	client := auth.New(configuration)

	return client.TokenRevoke()
}
