package main

import (
	"fmt"
	"time"

	config "github.com/chrusty/micro-pingpong/ping/config"
	handler "github.com/chrusty/micro-pingpong/ping/handler"
	player "github.com/chrusty/micro-pingpong/ping/player"
	proto "github.com/chrusty/micro-pingpong/ping/proto/src/go"
	logrus "github.com/sirupsen/logrus"

	microcli "github.com/micro/cli"
	microapi "github.com/micro/go-api"
	microgrpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
)

const (
	serviceName    = "com.cruft.service.ping"
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

			// Start playing pingpong:
			go player.Run(serviceConfig, service)

			return nil
		}),
	)

	// Register handler:
	proto.RegisterPingHandler(service.Server(), new(handler.PingHandler))

	// Register handler (REST):
	proto.RegisterPingHandler(service.Server(), new(handler.PingHandler), microapi.WithEndpoint(&microapi.Endpoint{
		// The RPC method
		Name: "Ping.Ping",
		// The HTTP paths. This can be a POSIX regex
		Path: []string{"/ping"},
		// The HTTP Methods for this endpoint
		Method: []string{"GET", "POST"},
		// The API handler to use
		Handler: microapi.Api,
	}))

	// Say something:
	logger.Debug("Starting instance")

	// Run the server:
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
