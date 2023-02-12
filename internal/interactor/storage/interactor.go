package storage

import (
    "github.com/hjhussaini/storage-srv-go/internal/contract"
)

type Interactor struct {
    store   contract.CloudStorage
}

func New(store contract.CloudStorage) *Interactor {
    return &Interactor{
        store:  store,
    }
}
