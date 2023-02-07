package v1

import (
    "net/http"

    "github.com/hjhussaini/storage-srv-go/internal/contract"

    "github.com/gorilla/mux"
)

type Handler struct {
    interactor  contract.StorageInteractor
}

func New(interactor contract.StorageInteractor) http.Handler {
    router := mux.NewRouter()
    handler := Handler{
        interactor: interactor,
    }

    router.HandleFunc("/scan/{directory}", handler.scanDirectory).Methods(http.MethodPost)
    router.HandleFunc("/storage/upload", handler.upload).Methods(http.MethodPost)

    return router
}
