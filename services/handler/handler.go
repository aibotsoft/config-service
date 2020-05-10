package handler

import (
	"context"
	"github.com/aibotsoft/config-service/pkg/store"
	pb "github.com/aibotsoft/gen/confpb"
	"github.com/aibotsoft/micro/config"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Handler struct {
	cfg *config.Config
	log *zap.SugaredLogger
	//api        *api.Api
	store *store.Store
}

func NewHandler(cfg *config.Config, log *zap.SugaredLogger, store *store.Store) *Handler {
	return &Handler{cfg: cfg, log: log, store: store}
}

func (h *Handler) Close() {
	h.store.Close()
}

func (h *Handler) GetConfig(ctx context.Context, serviceName string) (pb.ServiceConfig, error) {
	grpcPort, err := h.store.GetPortByName(ctx, serviceName)
	if err != nil {
		return pb.ServiceConfig{}, errors.Wrapf(err, "GetPortByName error for name: %q", serviceName)
	}
	return pb.ServiceConfig{GrpcPort: grpcPort}, nil

}

func (h *Handler) GetAccount(ctx context.Context, name string) (pb.Account, error) {
	return h.store.GetAccountByName(ctx, name)

}
func (h *Handler) GetCurrency(ctx context.Context) ([]pb.Currency, error) {
	return h.store.GetCurrency(ctx)
}
