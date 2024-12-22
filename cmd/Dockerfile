# Use the official Golang image as the base image
FROM golang:1.22 as builder

# Set the current working directory
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o cinema-service ./cmd/server.go

# Create a smaller image with the built app
FROM alpine:latest  

# Install necessary dependencies (like libpq for PostgreSQL connection)
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder image
COPY --from=builder /app/cinema-service .

# Expose port for gRPC server
EXPOSE 8080

# Run the service
CMD ["./cinema-service"]
