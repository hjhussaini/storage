package handlers

import (
	"net"
	"os"
	"time"

	"github.com/hjhussaini/storage/dbx"
	"github.com/hjhussaini/storage/models"
)

func (storage *Storage) uploadToDropbox(upload *models.Upload, save bool) {
	_, err := net.DialTimeout("tcp", "dropbox.com:80", 200*time.Millisecond)
	if err != nil {
		storage.logger.Printf("FAILED %s", err.Error())
		if save {
			upload.Save(storage.database)
		}
		return
	}

	err = dbx.Upload(
		*storage.dropboxConfiguration,
		upload.Source,
		upload.Destination,
	)
	if err != nil {
		storage.logger.Printf("FAILED %s", err.Error())
		if save {
			upload.Save(storage.database)
		}
		return
	}

	storage.logger.Printf("Upload '%s' successfully", upload.Source)
	if upload.RemoveSource {
		err = os.Remove(upload.Source)
		if err != nil {
			storage.logger.Printf("FAILED %s", err.Error())
			return
		}

		storage.logger.Printf("Remove '%s' successfully", upload.Source)
	}
}
