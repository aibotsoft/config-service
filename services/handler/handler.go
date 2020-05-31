package handler

import (
	"context"
	"github.com/aibotsoft/config-service/pkg/pin_client"
	"github.com/aibotsoft/config-service/pkg/store"
	pb "github.com/aibotsoft/gen/confpb"
	"github.com/aibotsoft/micro/config"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	cfg       *config.Config
	log       *zap.SugaredLogger
	store     *store.Store
	pinClient *pin_client.Client
	client    *http.Client
	NetStatus bool
}

func NewHandler(cfg *config.Config, log *zap.SugaredLogger, store *store.Store) *Handler {
	return &Handler{cfg: cfg, log: log, store: store, pinClient: pin_client.NewClient(cfg, log), client: &http.Client{}}
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
	//return h.store.GetAccountByName(ctx, name)
	return h.store.GetAccountByServiceName(ctx, name)

}
func (h *Handler) GetCurrency(ctx context.Context) ([]pb.Currency, error) {
	return h.store.GetCurrency(ctx)
}

func (h *Handler) GetServices(ctx context.Context) ([]pb.BetService, error) {
	return h.store.GetServices(ctx)
}
