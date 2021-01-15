package handlers

import (
	"net/http"

	"gitlab.com/hajihussaini/storage-service/dbx"
	"gitlab.com/hajihussaini/storage-service/models"
)

// authorized checks the request is authorized and calls next if ok
func (storage *Storage) authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if storage.dropboxConfiguration == nil {
			login := &models.Login{
				Storage: "dropbox",
			}
			login.Fetch(storage.database)
			if login.Token == "" {
				writer.WriteHeader(http.StatusNonAuthoritativeInfo)
				models.WriteJSON(
				GenericError{Message: "Non-Authoritative Information"},
				writer,
			)
				return
			}
			storage.dropboxConfiguration = dbx.Root(login.Token)
		}

		// Call the next handler, which can be another middleware in the chain,
		// or the final handler.
		next.ServeHTTP(writer, request)
	})
}
