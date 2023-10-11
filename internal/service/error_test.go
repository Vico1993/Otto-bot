package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnError(t *testing.T) {
	// Override templates
	templates = []string{"Super Error Messages"}

	res := ReturnError()
	assert.Equal(
		t,
		"Super Error Messages",
		res,
		"The template is not matching the expected text!",
	)
}
