package storage

import (
    "github.com/hjhussaini/storage-srv-go/internal/contract"
)

type Interactor struct {
    cloud   contract.CloudStorage
}

func New(cloud contract.CloudStorage) *Interactor {
    return &Interactor{
        cloud:  cloud,
    }
}
