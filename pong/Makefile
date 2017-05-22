all: protos build

protos:
	@mkdir -p proto/src/go
	@protoc -I proto --go_out=plugins=micro:proto/src/go proto/*.proto

build:
	@go build -o com.cruft.service.pong
