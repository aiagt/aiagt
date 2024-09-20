package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHandlers(t *testing.T) {
	_, err := ParseHandlers("../../app/user", "user")
	assert.NoError(t, err)
}
