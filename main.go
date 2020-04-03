package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab.com/hajihussaini/uploader-service/handlers"
	"gitlab.com/hajihussaini/uploader-service/models"
)

func main() {
	logFile, err := os.Create("/tmp/uploader.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "uploader-service", log.LstdFlags)
	validation := models.NewValidation()
	database, err := gorm.Open("sqlite3", "/tmp/uploader.sqlite3")
	if err != nil {
		logger.Fatal(err)
	}
	defer database.Close()

	router := mux.NewRouter()

	storageHandlers := handlers.NewStorage(logger, validation, database)
	storageHandlers.Register(router)
	storageHandlers.Migrate()

	server := http.Server{
		Addr:         ":2022",
		Handler:      router,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()

	killSignal := make(chan os.Signal)
	signal.Notify(killSignal, os.Interrupt)
	signal.Notify(killSignal, os.Kill)

	signaled := <-killSignal
	logger.Printf("Exit by '%s' signal", signaled)

	// Gracefully shut down the server
	timeout, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeout)
}
