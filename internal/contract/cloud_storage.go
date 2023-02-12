package contract

import (
    "os"
)

type CloudStorage interface {
    UploadFile(*os.File, string) error
}
