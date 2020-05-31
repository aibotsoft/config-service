package handler

import (
	"context"
	"net/http"
	"time"
)

const CheckNetUrl = "https://clients3.google.com/generate_204"
const CheckInetJobPeriod = time.Second * 30
const CheckInetTimeOut = time.Second * 10

func (h *Handler) CheckInetJob() {
	for {
		start := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), CheckInetTimeOut)
		err := h.CheckInetRound(ctx)
		cancel()
		if err != nil {
			h.log.Error(err)
		} else {
			h.log.Debugw("CheckInetJob_done", "time", time.Since(start))
		}
		time.Sleep(CheckInetJobPeriod)
	}
}
func (h *Handler) CheckInetRound(ctx context.Context) error {
	status, err := h.CheckInet(ctx)
	if err != nil {
		h.NetStatus = false
		return err
	}
	h.NetStatus = status
	return nil
}
func (h *Handler) CheckInet(ctx context.Context) (bool, error) {
	req, _ := http.NewRequestWithContext(ctx, "", CheckNetUrl, nil)
	resp, err := h.client.Do(req)
	if err != nil {
		return false, err
	}
	err = resp.Body.Close()
	return resp.StatusCode == 204, nil
}
