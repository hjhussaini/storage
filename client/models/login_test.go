package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingToken(t *testing.T) {
	login := Login{}
	validation := NewValidation()
	errs := validation.Validate(login)

	assert.Len(t, errs, 1)
}
