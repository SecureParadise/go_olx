.PHONY:build run pbuild
build:
	@go build -o bin/api ./cmd/api
run: build
	@./bin/api
pbuild:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o bin/api .cmd/api
