package main

import (
    "context"
    "log"

    "github.com/hjhussaini/storage-srv-go/config"
)

func main() {
    ctx := context.Background()
    cfg, err := config.NewFromEnv(ctx)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("running Storage server")
    log.Printf("listening on %s (HTTP)", cfg.HTTPServer.Address)
}
