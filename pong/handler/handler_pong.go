package handler

import (
	protopong "github.com/chrusty/micro-pingpong/pong/proto/src/go"

	logrus "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
)

type PongHandler struct{}

// Create a scheduled job:
func (pongHandler *PongHandler) Pong(ctx context.Context, req *protopong.PongRequest, rsp *protopong.PongResponse) error {

	rsp.Iteration = req.Iteration + 1
	rsp.Noise = "pong"

	logger.WithFields(logrus.Fields{"req_iteration": req.Iteration}).Info("Pong")

	return nil
}
