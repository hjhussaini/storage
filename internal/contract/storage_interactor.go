package contract

import (
    "context"

    "github.com/hjhussaini/storage-srv-go/internal/dto"
)

type StorageInteractor interface {
    Upload(context.Context, *dto.UploadRequest)
}
