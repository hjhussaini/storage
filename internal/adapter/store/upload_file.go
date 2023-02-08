package store

import (
    "io"
    "os"
    "time"

    "github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

const chunkSize int64 = 1 << 24

func (cloud *Cloud) UploadFile(file *os.File, path string) error {
    fileInfo, err := file.Stat()
    if err != nil {
        return err
    }

    fileSize := fileInfo.Size()
    now := time.Now().UTC().Round(time.Second)

    commitInfo := files.NewCommitInfo(path)
    commitInfo.Mode.Tag = "overwrite"
    commitInfo.ClientModified =&now 

    client := files.New(cloud.config)

    if fileSize > chunkSize {
        return uploadChunks(client, file, fileSize, commitInfo)
    } else {
        _, err = client.Upload(
            &files.UploadArg{
                CommitInfo: *commitInfo,
            },
            file,
        )
    }

    return err
}

func uploadChunks(
    client files.Client,
    content io.Reader,
    fileSize int64,
    commitInfo *files.CommitInfo,
) error {
    result, err := client.UploadSessionStart(
        &files.UploadSessionStartArg{},
        &io.LimitedReader{
            R:  content,
            N:  chunkSize,
        },
    )
    if err != nil {
        return err
    }

    uploadedBytes := chunkSize
    for (fileSize - uploadedBytes) > chunkSize {
        cursor := files.NewUploadSessionCursor(result.SessionId, uint64(uploadedBytes))
        arguments := files.NewUploadSessionAppendArg(cursor)
        err = client.UploadSessionAppendV2(
            arguments,
            &io.LimitedReader{
                R:  content,
                N:  chunkSize,
            },
        )
        if err != nil {
            return err
        }

        uploadedBytes += chunkSize
    }

    cursor := files.NewUploadSessionCursor(result.SessionId, uint64(uploadedBytes))
    arguments := files.NewUploadSessionFinishArg(cursor, commitInfo)
    _, err = client.UploadSessionFinish(arguments, content)

    return err
}
