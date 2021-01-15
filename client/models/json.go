package models

import (
	"encoding/json"
	"io"
)

// ReadJSON deserializes the object from JSON string
func ReadJSON(object interface{}, reader io.Reader) error {
	return json.NewDecoder(reader).Decode(object)
}

// WriteJSON serializes the given object into a string based JSON format
func WriteJSON(object interface{}, writer io.Writer) error {
	return json.NewEncoder(writer).Encode(object)
}
