# Micro PingPong

This is an example of how to use GoMicro's GRPC client & server to play RCP "PingPong"


## Instructions
* Install Go-Micro's protobuf compiler plugin: `go get github.com/micro/protobuf/protoc-gen-go`
* Run Consul (the default registry, easiest for local development): `docker run -d --net=bridge -p 8500:8500 --name consul consul agent -client=0.0.0.0  -server -bootstrap-expect 1 -data-dir=/tmp/consul`
* Build ping and pong (`make all` from each directory)
* Run ping: `./com.cruft.service.ping`
* Run pong: `./com.cruft.service.pong`
* Build a version of the "micro" binary with the GRPC client included (`_ "github.com/micro/go-plugins/client/grpc"`)


## Querying the ping and pong services with the "Micro Web" interface
* Run the "Micro Web" tool: `micro --client=grpc web`
* Explore the [registry](http://localhost:8082/registry)
* Use the [Query tool](http://localhost:8082/query) to query com.cruft.service.ping or com.cruft.service.pong


## Querying the ping and pong services using the Micro CLI tools
* List running services: `micro --client=grpc list services`
* Check the stats of ping: `micro --client=grpc stats com.cruft.service.ping`
* Query pong: `micro --client=grpc query com.cruft.service.pong Pong.Pong '{"iteration": 1}'`


## Querying the ping and pong services using CURL:
* Run the "Micro API" tool: `micro --enable_stats --client=grpc api --namespace=com.cruft.service`
* Query the pong service: `curl 'http://localhost:8080/pong/pong/pong?iteration=0'`
* Watch the live stats in the [API stats interface](http://localhost:8080/stats)
