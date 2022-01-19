FROM golang:1.17.6-alpine3.15 as builder
WORKDIR /build

ENV GOPROXY https://goproxy.cn
RUN ls -a
COPY . .
RUN go mod tidy
RUN go mod download



RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o mini-ecs ./cmd/main.go

FROM alpine:3.15 AS final
WORKDIR /app
COPY --from=builder /build/mini-ecs /app/
COPY --from=builder /build/configs /app/configs

ENTRYPOINT ["/app/mini-ecs"]