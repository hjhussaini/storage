package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingSource(t *testing.T) {
	upload := Upload{
		Destination: "/dst",
	}
	validation := NewValidation()
	errs := validation.Validate(upload)

	assert.Len(t, errs, 1)
}

func TestMissingDestination(t *testing.T) {
	upload := Upload{
		Source: "/tmp/test.txt",
	}
	validation := NewValidation()
	errs := validation.Validate(upload)

	assert.Len(t, errs, 1)
}

func TestPathMissingPrefix(t *testing.T) {
	upload := Upload{
		Source:      "/tmp/test.txt",
		Destination: "path",
	}
	validation := NewValidation()
	errs := validation.Validate(upload)

	assert.Len(t, errs, 1)
}

func TestPathHasSuffix(t *testing.T) {
	upload := Upload{
		Source:      "/tmp/test.txt",
		Destination: "/path/",
	}
	validation := NewValidation()
	errs := validation.Validate(upload)

	assert.Len(t, errs, 1)
}
