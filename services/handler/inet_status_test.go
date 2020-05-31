package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_CheckInet(t *testing.T) {
	_, err := h.CheckInet(context.Background())
	assert.NoError(t, err)
}

func TestHandler_CheckInetRound(t *testing.T) {
	err := h.CheckInetRound(context.Background())
	assert.NoError(t, err)
	err = h.CheckInetRound(context.Background())
	err = h.CheckInetRound(context.Background())
}
