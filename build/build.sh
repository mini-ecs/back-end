#!/bin/sh
echo "============================  generating swagger api  ==========================="
swag init -generalInfo cmd/main.go -d ./
echo ""
echo "============================  formatting swagger api  ==========================="
swag fmt -g cmd/main.go -d ./
echo ""
echo "===============================  starting server  ==============================="


export PATH=$PATH:/home/fangaoyang/work/noVNC-1.3.0/utils
go run cmd/main.go