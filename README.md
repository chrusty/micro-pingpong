# Micro PingPong

This is an example of how to use GoMicro's GRPC client & server to play RCP "PingPong"

## Instructions
* Install Go-Micro's protobuf compiler plugin: `go get github.com/micro/protobuf/protoc-gen-go`
* Build ping and pong (`make all` from each directory)
* Run ping: `./com.cruft.service.ping`
* Run pong: `./com.cruft.service.pong`

## Todo
* A demo of how to route REST calls through the Micro API
