FROM golang:1.24-alpine

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Copy the .env file
COPY .env /app/.env

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o ordersystem ./cmd/ordersystem

# Expose ports for web server, gRPC server, and GraphQL server
EXPOSE 8000 50051 8080

# Set the entry point
CMD ["/app/ordersystem"]
