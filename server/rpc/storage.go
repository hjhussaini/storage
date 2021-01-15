package rpc

import (
	"github.com/hjhussaini/storage/proto"
)

type Storage struct{}

func (storage *Storage) Upload(stream proto.Storage_UploadServer) error {

	return nil
}
