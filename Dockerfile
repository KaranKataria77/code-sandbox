# Use Ubuntu as the base image
FROM ubuntu:latest

# Set non-interactive mode (prevents prompts during package installation)
ENV DEBIAN_FRONTEND=noninteractive

# Install required programming languages (Python, Node.js, Go)
RUN apt-get update && apt-get install -y \
    bash \
    curl \
    wget \
    net-tools \
    python3 python3-pip \
    nodejs npm \
    golang \
    protobuf-compiler \
    && apt-get clean


# Install gRPC go dependency
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Create a non-root user for security
RUN useradd -m sandboxuser
WORKDIR /home/sandboxuser

# copy gRPC server into container
COPY --chown=sandboxuser:sandboxuser . .

# Build gRPC server
RUN go mod tidy && go build -o grpc-server ./server.go

# Expose gRPC port
EXPOSE 50051

# Switch to non-root user
USER sandboxuser

# Default command (can be modified later)
CMD ["/home/sandboxuser/grpc-server"]
