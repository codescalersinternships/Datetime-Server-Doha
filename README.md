# Date Server 
Create a basic HTTP server that returns the current date and time.

## Implementation
- standards net/http package
- gin framework


## Installation 
- 1. Clone project
```golang
   https://github.com/codescalersinternships/Datetime-Server-Doha.git
```
- 2. Run server by 2 ways :
  - docker compose
  - makefile commands

### Docker-compose 
```golang
docker-compose up // to build the docker image/containers
docker-compose down // to stop the image/container
```
### Makefile
 - you can build binaries directly
```golang
make build // will build binary files for the servers
./http // will run server implemented uding http
./gin  // will run server implemented using gin framework
```
 - build images of docker then containers
``` golang
make build-gin-image // build gin server image 
build-gin-container // build gin server container

build-http-image // build http server image
build-http-container // build http server container
```
## test
- to run all tests
```golang
  go test v ./...
```
## format
- format all files inside project
```golang
gofmt -w .
```
