linter:
	golangci-lint run ./...

build:
	go build -o ./http ./cmd/gin/main.go
	go build -o ./gin ./cmd/net-http/main.go

format:
	gofmt -w .


test:
	go test -v ./...
