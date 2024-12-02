# go-microservices-grpc-mongo

This project consists of an API Gateway (REST) and an Auth Service (gRPC).

## Architecture Overview

- **API Gateway**: Acts as the main entry point for client applications, exposing REST endpoints and handling routing, validation, and authentication.
- **Auth Service**: A microservice built with gRPC, responsible for user authentication and token management.

---

## Features

### API Gateway

- RESTful endpoints for client interaction.
- Handles authentication and forwards authorized requests to appropriate microservices.
- Validation of input payloads.
- Centralized error handling and response formatting.

### Auth Service

- Implements gRPC for efficient inter-service communication.
- User authentication (login and token validation).
- Token generation and management (e.g., JWT).

---

## How to run

### run prod container

- `docker compose -f docker-compose-prod up`

### run dev container

- `docker compose up .`

#### exec container

- `docker exec -it go-microservices-api-gateway bash`
- `docker exec -it go-microservices-auth-svc bash`

#### how to Generate pb

- `protoc --go_out=. --go-grpc_out=. ./pkg/messages/*.proto`
