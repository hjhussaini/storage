package storage

import (
    "context"
    "log"
    "os"

    "github.com/hjhussaini/storage-srv-go/internal/dto"
)

func (interactor *Interactor) UploadFile(ctx context.Context, request *dto.UploadRequest) error {
    file, err := os.Open(request.File)
    if err != nil {
        log.Println(err)

        return err
    }

    log.Println("uploading", request.File)
    err = interactor.cloud.UploadFile(file, request.To)
    file.Close()
    if err != nil {
        log.Printf("could not upload '%s': %v", request.File, err)

        return err
    }

    log.Println("uploaded", request.File)
    if request.DeleteFile {
        return os.Remove(request.File)
    }

    return nil
}
