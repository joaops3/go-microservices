## Generate pb

`protoc --go_out=. --go-grpc_out=. ./pkg/messages/*.proto`

## run container

`docker exec -it go-microservices-api-gateway bash`
