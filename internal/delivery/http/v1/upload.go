package v1

import (
    "encoding/json"
    "net/http"

    "github.com/hjhussaini/storage-srv-go/internal/dto"
)

func (handler *Handler) upload(response http.ResponseWriter, request *http.Request) {
    var uploadRequest   dto.UploadRequest

    response.Header().Set("Content-Type", "application/json")
    err := json.NewDecoder(request.Body).Decode(&uploadRequest)
    request.Body.Close()
    if err != nil {
        response.WriteHeader(http.StatusBadRequest)

        return
    }

    handler.interactor.Upload(request.Context(), &uploadRequest)
    response.WriteHeader(http.StatusOK)
}
