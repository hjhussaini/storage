package dbx

import (
	"io"
	"os"
	"time"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/mitchellh/ioprogress"
)

const chunkSize int64 = 1 << 24

func Upload(
	configuration dropbox.Config,
	sourcePath string,
	destinationPath string,
) error {
	contents, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer contents.Close()

	contentsInformation, err := contents.Stat()
	if err != nil {
		return err
	}

	progressbar := &ioprogress.Reader{
		Reader: contents,
		Size:   contentsInformation.Size(),
	}

	commitInformation := files.NewCommitInfo(destinationPath)
	commitInformation.Mode.Tag = "overwrite"
	// The DropBox API only accepts timestamps in UTC with second precission.
	commitInformation.ClientModified = time.Now().UTC().Round(time.Second)

	client := files.New(configuration)
	if contentsInformation.Size() > chunkSize {
		return uploadChunked(
			client,
			progressbar,
			commitInformation,
			contentsInformation.Size(),
		)
	}

	if _, err = client.Upload(commitInformation, progressbar); err != nil {
		return err
	}

	return nil
}

func uploadChunked(
	client files.Client,
	reader io.Reader,
	commitInformation *files.CommitInfo,
	totalSize int64,
) error {
	result, err := client.UploadSessionStart(
		files.NewUploadSessionStartArg(),
		&io.LimitedReader{R: reader, N: chunkSize},
	)
	if err != nil {
		return err
	}

	written := chunkSize

	for (totalSize - written) > chunkSize {
		cursor := files.NewUploadSessionCursor(result.SessionId, uint64(written))
		arguments := files.NewUploadSessionAppendArg(cursor)

		err = client.UploadSessionAppendV2(
			arguments,
			&io.LimitedReader{R: reader, N: chunkSize},
		)
		if err != nil {
			return err
		}
		written += chunkSize
	}

	cursor := files.NewUploadSessionCursor(result.SessionId, uint64(written))
	arguments := files.NewUploadSessionFinishArg(cursor, commitInformation)

	_, err = client.UploadSessionFinish(arguments, reader)
	if err != nil {
		return err
	}

	return nil
}
