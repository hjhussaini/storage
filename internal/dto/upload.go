package dto

// swagger:model
type UploadRequest struct {
    // The file to be uploaded
    // required: true
    File        string  `json:"file"`
    // The path to upload file
    // required: true
    To          string  `json:"to"`
    // Determines to delete file after uploading or not
    // required: false
    DeleteFile  bool    `json:"delete_file,omitempty"`
}
