#!/bin/sh
echo "============================  generating swagger api  ==========================="
swag init -generalInfo cmd/main.go -d ./
echo ""
echo "============================  formatting swagger api  ==========================="
swag fmt -g cmd/main.go -d ./
echo ""
echo "===============================  starting server  ==============================="
sleep 3
go run cmd/main.go