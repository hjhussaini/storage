package store

import (
    "github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
)

// Cloud implements contract.StorageCloud interface
type Cloud struct {
    config  dropbox.Config
}

func New(token string) *Cloud {
    config := dropbox.Config{
        Token:      token,
        LogLevel:   dropbox.LogOff,
    }

    return &Cloud{
        config: config,
    }
}
