package rpc

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/hjhussaini/storage/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Storage struct {
	Path string
}

func (storage *Storage) Upload(stream proto.Storage_UploadServer) error {
	request, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "Could not receive data: %s", err)
	}

	if request.GetMachineId() == 0 {
		return status.Errorf(codes.Internal, "Invalid machine ID")
	}

	path := fmt.Sprintf("%s/%d/", storage.Path, request.GetMachineId())
	err = os.Mkdir(path, 0664)
	if err != nil {
		if !os.IsExist(err) {
			return status.Errorf(
				codes.Internal,
				"Could not create directory: %s", err.Error(),
			)
		}
	}

	imageFile := path + request.GetFileName()
	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		request, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return status.Errorf(codes.Internal, "Could not receive chunk data: %s", err)
		}

		chunk := request.GetChunk()
		imageSize += len(chunk)
		_, err = imageData.Write(chunk)
	}

	file, err := os.Create(imageFile)
	if err != nil {
		return status.Errorf(codes.Internal, "Could not save image: %s", err)
	}
	defer file.Close()

	_, err = imageData.WriteTo(file)
	if err != nil {
		return status.Errorf(codes.Internal, "Could not write to file: %s", err)
	}

	response := &proto.FileResponse{
		Size_: uint32(imageSize),
	}
	err = stream.SendAndClose(response)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to send response: %s", err)
	}

	return nil
}
