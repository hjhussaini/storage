package contract

import (
    "context"

    "github.com/hjhussaini/storage-srv-go/internal/dto"
)

type StorageInteractor interface {
    UploadFile(context.Context, *dto.UploadRequest) error
    ScanDirectory(context.Context, *dto.ScanRequest)
}
