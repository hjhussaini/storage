package dto

// swagger:model
type LoginRequest struct {
    // The application token
    // required: true
    Token   string  `json:"token"`
}
