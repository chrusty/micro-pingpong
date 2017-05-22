package handler

import (
	protoping "github.com/chrusty/micro-pingpong/ping/proto/src/go"

	logrus "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
)

type PingHandler struct{}

// Create a scheduled job:
func (pingHandler *PingHandler) Ping(ctx context.Context, req *protoping.PingRequest, rsp *protoping.PingResponse) error {

	rsp.Iteration = req.Iteration + 1
	rsp.Noise = "ping"

	logger.WithFields(logrus.Fields{"req_iteration": req.Iteration}).Info("Ping")

	return nil
}
