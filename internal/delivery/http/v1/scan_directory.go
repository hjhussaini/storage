package v1

import (
    "encoding/json"
    "net/http"

    "github.com/hjhussaini/storage-srv-go/internal/dto"
)

func (handler *Handler) scanDirectory(response http.ResponseWriter, request *http.Request) {
    var scanRequest dto.ScanRequest

    response.Header().Set("Content-Type", "application/json")
    err := json.NewDecoder(request.Body).Decode(&scanRequest)
    request.Body.Close()
    if err != nil {
        response.WriteHeader(http.StatusBadRequest)

        return
    }

    handler.interactor.ScanDirectory(request.Context(), &scanRequest)

    response.WriteHeader(http.StatusOK)
}
