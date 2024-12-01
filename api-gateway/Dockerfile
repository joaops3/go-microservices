# Use the official Golang image as the base image
FROM golang:1.22.7

# Set the Current Working Directory inside the container
WORKDIR /app
ENV PATH="/go/bin:${PATH}"

# Copy go mod and sum files
COPY ./go.mod ./go.sum ./

RUN export PATH="$PATH:/go/bin"

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install protobuf compiler
RUN apt-get update && apt-get install -y protobuf-compiler

# Install protoc-gen-go and protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest



# Expose port 3333 to the outside world
EXPOSE 3333

# Command to keep the container running
CMD ["tail", "-f", "/dev/null"]