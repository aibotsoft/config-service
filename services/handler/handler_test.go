package handler

import (
	"context"
	"github.com/aibotsoft/config-service/pkg/store"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/aibotsoft/micro/sqlserver"
	"github.com/stretchr/testify/assert"
	"testing"
)

var h *Handler

func TestMain(m *testing.M) {
	cfg := config.New()
	log := logger.New()
	db := sqlserver.MustConnectX(cfg)
	sto := store.NewStore(cfg, log, db)
	h = NewHandler(cfg, log, sto)
	m.Run()
	h.Close()
}

func TestHandler_GetConfig(t *testing.T) {
	got, err := h.GetConfig(context.Background(), "fuck-off")
	assert.Error(t, err)
	got, err = h.GetConfig(context.Background(), "config-service")
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got)
	}
}

func TestHandler_GetAccount(t *testing.T) {
	got, err := h.GetAccount(context.Background(), "fuck-off")
	assert.Error(t, err)
	got, err = h.GetAccount(context.Background(), "sbo-service")
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got)
	}
}

func TestHandler_GetCurrency(t *testing.T) {
	got, err := h.GetCurrency(context.Background())
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got)
	}
}

func TestHandler_GetServices(t *testing.T) {
	got, err := h.GetServices(context.Background())
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got)
	}
}
