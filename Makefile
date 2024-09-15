linter:
	golangci-lint run ./...

build:
	go build -o ./cmd/gin/ ./cmd/gin/main.go
	go build -o ./cmd/net-http/ ./cmd/net-http/main.go

format:
	gofmt -w .

