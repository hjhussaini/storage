package handlers

import (
	"net/http"

	"gitlab.com/hajihussaini/storage-service/models"
)

func (storage *Storage) validate(
	object interface{},
	writer http.ResponseWriter,
	request *http.Request,
) bool {
	writer.Header().Add("content-type", "application/json")

	err := models.ReadJSON(object, request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		models.WriteJSON(GenericError{Message: err.Error()}, writer)

		return false
	}

	// Validate the object
	errs := storage.validation.Validate(object)
	if len(errs) != 0 {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		models.WriteJSON(ValidationError{Messages: errs.Errors()}, writer)

		return false
	}

	return true
}
