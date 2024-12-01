#!/bin/bash

go clean --cache && go test -v -cover api-gateway...

go clean --cache && go test -v -cover auth-svc/...

go build -o ./api-gateway ./api-gateway/cmd/main.go
go build -o ./auth-svc ./auth-svc/cmd/main.go