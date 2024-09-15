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
- 2. Run server by 3 ways :
  - docker compose
  - makefile commands
  - importing the correct pkg url in your project 

### Docker-compose 
```golang
docker-compose up // to build the docker image/containers
docker-compose down // to stop the image/container
```
### Makefile
```golang
make build // will build binary files for the servers
./http // will run server implemented uding http
./gin  // will run server implemented using gin framework
```

### Import Package
- 1. server impelented by http
```golang
import pkghttp "github.com/dohaelsawy/codescalers/datetimeserver/pkg/net-http"
```
- 2. server impelented by gin
```golang
import pkgin "github.com/dohaelsawy/codescalers/datetimeserver/pkg/gin"
```
## test
- to run all tests
```golang
  go test v ./...
```
