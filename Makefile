	USER=fangaoyang
	HOST=10.249.46.250
	DIR=/home/fangaoyang/work/backend-sync
all: swag proto build

swag:
	swag init --parseDependency --parseInternal -generalInfo cmd/main.go -d ./
	swag fmt -g cmd/main.go -d ./

build:
	go build cmd/main.go

docker:
	docker build -f build/Dockerfile . -t miniecs/backend:v1

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/proto/libvirt.proto

sync:
	rsync -avz --delete ./ $(USER)@$(HOST):$(DIR)