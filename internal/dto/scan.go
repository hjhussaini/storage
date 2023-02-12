package dto

// swagger:model
type ScanRequest struct {
    // The directory path to scan
    // required: true
    Directory   string  `json:"directory"`
    // The path of upload
    // required: true
    To          string  `json:"to"`
    // Determines to delete file after uploading
    // required: false
    DeleteFile  bool    `json:"delete_file,omitempty"`
}
