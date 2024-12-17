package server

import (
	"go-microservices-grpc/api-gateway/internal/api/router"
	config "go-microservices-grpc/api-gateway/internal/configs"
)

type HttpServer struct {
}

func (s *HttpServer) Serve() error {
	logger := config.GetLogger("MAIN")
	config.LoadEnvFile()
	

	err := router.Initialize()
	if err != nil {
		logger.Errorf(err.Error())
		return err
	}

	return nil
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}