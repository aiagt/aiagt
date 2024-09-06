package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseHandlers(t *testing.T) {
	_, err := ParseHandlers("../../app/user", "user")
	assert.NoError(t, err)
}
