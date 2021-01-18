package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/hjhussaini/storage/handlers"
	"github.com/hjhussaini/storage/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	port := os.Getenv("PORT")
	databasePath := os.Getenv("DATABASE")
	logFilePath := os.Getenv("LOG_FILE")

	logFile, err := os.Create(logFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "storage", log.LstdFlags)
	validation := models.NewValidation()
	database, err := gorm.Open("sqlite3", databasePath)
	if err != nil {
		logger.Fatal(err)
	}
	defer database.Close()

	router := mux.NewRouter()

	storageHandlers := handlers.NewStorage(logger, validation, database)
	storageHandlers.Register(router)
	storageHandlers.Migrate()

	server := http.Server{
		Addr:         ":" + port,
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
