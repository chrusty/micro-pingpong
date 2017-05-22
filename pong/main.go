package main

import (
	"fmt"
	"time"

	config "github.com/chrusty/micro-pingpong/pong/config"
	handler "github.com/chrusty/micro-pingpong/pong/handler"
	proto "github.com/chrusty/micro-pingpong/pong/proto/src/go"
	logrus "github.com/sirupsen/logrus"

	microcli "github.com/micro/cli"
	microgrpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/transport/grpc"
)

const (
	serviceName    = "com.cruft.service.pong"
	serviceVersion = "0.0.1"
)

var (
	logger        = logrus.WithFields(logrus.Fields{"logger": "main"})
	serviceConfig = config.Config{}
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	// Create a new service. Optionally include some options here.
	service := microgrpc.NewService(
		micro.Name(serviceName),
		micro.Version(serviceVersion),
		micro.RegisterTTL(time.Second*11),
		micro.RegisterInterval(time.Second*5),

		// Add runtime flags:
		micro.Flags(
			microcli.DurationFlag{
				Name:        "hold_time",
				Value:       2 * time.Second,
				Usage:       "How long to hold the ball before returning it",
				EnvVar:      "HOLD_TIME",
				Destination: &serviceConfig.HoldTime,
			},
		),
	)

	// Init:
	service.Init(
		// Register a BeforeStart handler:
		micro.BeforeStart(func() error {

			// Initialise the handler:
			handler.Initialise(serviceConfig)

			return nil
		}),
	)

	// Register handler:
	proto.RegisterPongHandler(service.Server(), new(handler.PongHandler))

	// Say something:
	logger.Debug("Starting instance")

	// Run the server:
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
