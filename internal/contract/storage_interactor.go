package contract

import (
    "context"

    "github.com/hjhussaini/storage-srv-go/internal/dto"
)

type StorageInteractor interface {
    Login(context.Context, *dto.LoginRequest) error
    Upload(context.Context, *dto.UploadRequest) error
}
