lint:
	golangci-lint run ./...

build:
	go build -o ./http ./cmd/gin/main.go
	go build -o ./gin ./cmd/net-http/main.go

format:
	gofmt -w .


test:
	go test -v ./...


build-gin-image:
	docker build . -f Dockerfile.gin -t doha/gin

build-http-image:
	docker build . -f Dockerfile.http -t doha/http

build-gin-container:
	docker run -p 8080:8080 --name gincontainer doha/gin

build-http-container:
	docker run -p 8090:8090 --name httpcontainer2 doha/http



