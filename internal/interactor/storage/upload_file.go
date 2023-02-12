package storage

import (
    "context"
    "os"

    "github.com/hjhussaini/storage-srv-go/internal/dto"
)

func (interactor *Interactor) UploadFile(ctx context.Context, request *dto.UploadRequest) error {
    file, err := os.Open(request.File)
    if err != nil {
        return err
    }

    err = interactor.cloud.UploadFile(file, request.To)
    file.Close()
    if err != nil {
        return err
    }

    if request.DeleteFile {
        return os.Remove(request.File)
    }

    return nil
}
