package config

import (
    "context"
    "time"

    "github.com/sethvargo/go-envconfig"
)

// DropBoxConfig represents DropBox configuration
type DropBoxConfig struct {
    Token   string  `env:"DROPBOX_TOKEN,required"`
}

// HTTPServerConfig represents HTTP server configuration
type HTTPServerConfig struct {
    Address             string          `env:"HTTP_SERVER_ADDRESS,default=0.0.0.0:80"`
    ReadHeaderTimeout   time.Duration   `env:"HTTP_SERVER_READ_HEADER_TIMEOUT,default=5s"`
    ReadTimeout         time.Duration   `env:"HTTP_SERVER_READ_TIMEOUT,default=10s"`
    WriteTimeout        time.Duration   `env:"HTTP_SERVER_WRITE_TIMEOUT,default=10s"`
}

// Configuration represents service configuration
type Configuration struct {
    DropBox     DropBoxConfig
    HTTPServer  HTTPServerConfig
}

func NewFromEnv(ctx context.Context) (*Configuration, error) {
    var configuration Configuration

    err := envconfig.Process(ctx, &configuration)

    return &configuration, err
}
