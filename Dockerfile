# Use official Golang image as the base (updated to 1.23)
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o todo-api main.go

# Use a smaller base image for the final container
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/todo-api .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./todo-api"]