

all: swag proto build

swag:
	swag init --parseDependency --parseInternal -generalInfo cmd/main.go -d ./
	swag fmt -g cmd/main.go -d ./

build:
	go build cmd/main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/proto/libvirt.proto