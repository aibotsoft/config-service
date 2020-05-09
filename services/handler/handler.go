package handler

import (
	"github.com/aibotsoft/config-service/pkg/store"
	"github.com/aibotsoft/micro/config"
	"go.uber.org/zap"
)

type Handler struct {
	cfg *config.Config
	log *zap.SugaredLogger
	//api        *api.Api
	store *store.Store
}

func (h *Handler) Close() {
	h.store.Close()
}
func NewHandler(cfg *config.Config, log *zap.SugaredLogger, store *store.Store) *Handler {
	return &Handler{cfg: cfg, log: log, store: store}
}
