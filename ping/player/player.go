package player

import (
	"time"

	config "github.com/chrusty/micro-pingpong/ping/config"
	protopong "github.com/chrusty/micro-pingpong/pong/proto/src/go"

	micro "github.com/micro/go-micro"
	logrus "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
)

var (
	logger          = logrus.WithFields(logrus.Fields{"logger": "player"})
	iteration int32 = 0
)

func Run(serviceConfig config.Config, service micro.Service) {

	// Use the generated client in Pong's proto:
	pongClient := protopong.NewPongClient("com.cruft.service.pong", service.Client())

	for {
		// Sleep for a bit:
		time.Sleep(serviceConfig.HoldTime)

		// Make a Pong request:
		pongResponse, err := pongClient.Pong(context.Background(), &protopong.PongRequest{
			Iteration: iteration,
		})
		if err != nil {
			logger.WithFields(logrus.Fields{"iteration": iteration, "error": err}).Error("Couldn't talk to pong")
		} else {
			iteration = pongResponse.GetIteration()
			logger.WithFields(logrus.Fields{"iteration": iteration, "response": pongResponse}).Info("Ping")
		}
	}

}
