package storage

import (
    "context"
    "log"
    "os"

    "github.com/hjhussaini/storage-srv-go/internal/dto"
)

func (interactor *Interactor) ScanDirectory(ctx context.Context, request *dto.ScanRequest) {
    entries, err := os.ReadDir(request.Directory)
    if err != nil {
        log.Println(err)

        return
    }

    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }

        uploadFileRequest := &dto.UploadRequest{
            File:       request.Directory + "/" + entry.Name(),
            To:         request.To + entry.Name(),
            DeleteFile: request.DeleteFile,
        }
        go interactor.UploadFile(ctx, uploadFileRequest)
    }
}
