package pin_client

import (
	"context"
	api "github.com/aibotsoft/gen/pinapi"
	"github.com/aibotsoft/micro/config"
	"go.uber.org/zap"
)

type Client struct {
	cfg *config.Config
	log *zap.SugaredLogger
	*api.APIClient
}

func (c *Client) GetCurrencies(ctx context.Context) ([]api.Currency, error) {
	resp, _, err := c.OthersApi.CurrenciesV2Get(ctx).Execute()
	if err != nil {
		return nil, err
	}
	return resp.GetCurrencies(), nil
}

func NewClient(cfg *config.Config, log *zap.SugaredLogger) *Client {
	clientConfig := api.NewConfiguration()
	clientConfig.Debug = cfg.Service.Debug
	return &Client{cfg: cfg, log: log, APIClient: api.NewAPIClient(clientConfig)}
}
