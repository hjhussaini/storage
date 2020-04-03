package handlers

import (
	"net/http"

	"gitlab.com/hajihussaini/uploader-service/dbx"
	"gitlab.com/hajihussaini/uploader-service/models"
)

// swagger:route POST /dropbox/login
func (storage *Storage) dropbox_login(
	writer http.ResponseWriter,
	request *http.Request) {
	login := models.Login{
		Storage: "dropbox",
	}
	if !storage.validate(&login, writer, request) {
		return
	}
	storage.dropboxConfiguration = dbx.Root(login.Token)
	login.Save(storage.database)
}

// swagger:route POST /dropbox/upload
func (storage *Storage) dropbox_upload(
	writer http.ResponseWriter,
	request *http.Request) {
	upload := models.Upload{
		Storage: "dropbox",
	}
	if !storage.validate(&upload, writer, request) {
		return
	}

	go storage.uploadToDropbox(&upload, true)
}

// swagger:route POST /dropbox/scan
func (storage *Storage) dropbox_scan(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var uploads models.Uploads

	uploads.Fetch(storage.database, "dropbox")
	total := len(uploads)

	for index, upload := range uploads {
		storage.logger.Printf("(%d of %d) '%s'", index+1, total, upload.Source)
		storage.uploadToDropbox(&upload, false)
	}
}

// swagger:route POST /dropbox/logout
func (storage *Storage) dropbox_logout(
	writer http.ResponseWriter,
	request *http.Request) {
	err := dbx.Logout(*storage.dropboxConfiguration)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		models.WriteJSON(GenericError{Message: err.Error()}, writer)
		return
	}
}
