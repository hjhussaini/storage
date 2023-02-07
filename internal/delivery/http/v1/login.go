package v1

import (
    "encoding/json"
    "net/http"

    "github.com/hjhussaini/storage-srv-go/internal/dto"
)

func (handler *Handler) login(response http.ResponseWriter, request *http.Request) {
    var loginRequest    dto.LoginRequest

    response.Header().Set("Content-Type", "application/json")
    err := json.NewDecoder(request.Body).Decode(&loginRequest)
    request.Body.Close()
    if err != nil {
        response.WriteHeader(http.StatusBadRequest)

        return
    }

    err = handler.interactor.Login(request.Context(), &loginRequest)
    if err != nil {
        response.WriteHeader(http.StatusInternalServerError)

        return
    }

    response.WriteHeader(http.StatusOK)
}
