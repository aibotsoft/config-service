package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_CollectCurrency(t *testing.T) {
	got, err := h.CollectCurrency(context.Background())
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got)
	}
}

func TestHandler_CurrencyRound(t *testing.T) {
	err := h.CurrencyRound(context.Background())
	if assert.NoError(t, err) {
	}
}
