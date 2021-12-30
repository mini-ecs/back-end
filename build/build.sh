#!/bin/sh

swag init -generalInfo cmd/main.go -d ./
swag fmt -g cmd/main.go -d ./