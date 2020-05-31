package handler

import (
	"context"
	api "github.com/aibotsoft/gen/pinapi"
)

func (h *Handler) CurrencyRound(ctx context.Context) error {
	resp, err := h.CollectCurrency(ctx)
	if err != nil {
		return err
	}
	err = h.store.SaveCurrency(ctx, resp)
	return err
}
func (h *Handler) CollectCurrency(ctx context.Context) ([]api.Currency, error) {
	account, err := h.GetAccount(ctx, "pin-service")
	if err != nil {
		return nil, err
	}
	auth := context.WithValue(ctx, api.ContextBasicAuth, api.BasicAuth{UserName: account.Username, Password: account.Password})
	resp, err := h.pinClient.GetCurrencies(auth)
	return resp, err
	//return h.store.GetCurrency(ctx)
}
