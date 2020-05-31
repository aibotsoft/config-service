package handler

import (
	"context"
	api "github.com/aibotsoft/gen/pinapi"
	"time"
)

const CurrencyJobPeriod = time.Hour

func (h *Handler) CurrencyJob() {
	for {
		start := time.Now()
		if !h.NetStatus {
			h.log.Info("netStatus_not_ok")
			time.Sleep(time.Minute)
			continue
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := h.CurrencyRound(ctx)
		cancel()
		if err != nil {
			h.log.Error(err)
		} else {
			h.log.Debugw("CurrencyJob_done", "time", time.Since(start))
		}
		time.Sleep(CurrencyJobPeriod)
	}
}
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
}
