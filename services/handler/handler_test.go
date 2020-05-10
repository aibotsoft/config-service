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

func InitHelper(t *testing.T) *Handler {
	t.Helper()
	cfg := config.New()
	log := logger.New()
	db := sqlserver.MustConnectX(cfg)
	sto := store.NewStore(cfg, log, db)
	h := NewHandler(cfg, log, sto)
	return h
}

func TestHandler_GetConfig(t *testing.T) {
	h := InitHelper(t)
	got, err := h.GetConfig(context.Background(), "fuck-off")
	assert.Error(t, err)
	got, err = h.GetConfig(context.Background(), "config-service")
	if assert.NoError(t, err) {
		assert.NotEmpty(t, got)
	}
}
