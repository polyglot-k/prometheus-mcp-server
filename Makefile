.PHONY: run build docker-build

run:
	go run cmd/main.go

build:
	go build -o prometheus-mcp.exe cmd/main.go

docker-build:
	docker build -t prometheus-mcp .
