package handlers

import (
	"log"
	"net/http"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gitlab.com/hajihussaini/uploader-service/models"
)

type Storage struct {
	logger               *log.Logger
	validation           *models.Validation
	database             *gorm.DB
	dropboxConfiguration *dropbox.Config
}

func (storage *Storage) Register(router *mux.Router) {
	loginRouter := router.Methods(http.MethodPost).Subrouter()
	loginRouter.HandleFunc("/dropbox/login", storage.dropbox_login)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.Use(storage.authorized)
	postRouter.HandleFunc("/dropbox/upload", storage.dropbox_upload)
	postRouter.HandleFunc("/dropbox/scan", storage.dropbox_scan)
	postRouter.HandleFunc("/dropbox/logout", storage.dropbox_logout)
}

func (storage *Storage) Migrate() {
	storage.database.SingularTable(true)

	storage.database.AutoMigrate(&models.Login{})
	storage.database.AutoMigrate(models.Upload{})
}

func NewStorage(
	logger *log.Logger,
	validation *models.Validation,
	database *gorm.DB,
) *Storage {
	return &Storage{logger, validation, database, nil}
}
