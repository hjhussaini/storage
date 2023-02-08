package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    "log"

    "github.com/hjhussaini/storage-srv-go/config"
    "github.com/hjhussaini/storage-srv-go/internal/adapter/store"
    "github.com/hjhussaini/storage-srv-go/internal/delivery/http/v1"
)

func main() {
    ctx := context.Background()
    cfg, err := config.NewFromEnv(ctx)
    if err != nil {
        log.Fatal(err)
    }

    store.New(cfg.DropBox.Token)

    log.Println("running Storage server")

    httpHandler := v1.New(nil)
    errs := make(chan error, 2)
    server := http.Server{
        Addr:               cfg.HTTPServer.Address,
        Handler:            httpHandler,
        ReadHeaderTimeout:  cfg.HTTPServer.ReadHeaderTimeout,
        ReadTimeout:        cfg.HTTPServer.ReadTimeout,
        WriteTimeout:       cfg.HTTPServer.WriteTimeout,
    }

    go func() {
        if err := server.ListenAndServe(); err != nil {
            errs <- fmt.Errorf("could not serve server: %s", err.Error())
        }

        errs <- nil
    }()
    log.Printf("listening on %s (HTTP)", cfg.HTTPServer.Address)

    go func() {
        stop := make(chan os.Signal)
        signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
        <-stop

        // gracefully shut down the server
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        server.Shutdown(ctx)
        log.Println("shut down Storage server gracefully")
    }()

    if err := <-errs; err != nil {
        log.Fatal(err)
    }
}
