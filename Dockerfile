# Build stage
FROM golang:1.23 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application (replace `taurus` with your binary name)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./taurus ./cmd/serveHttp

# Final stage: minimal image
FROM alpine:latest

# Set up a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/taurus ./taurus

# Change ownership to the non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose the port the service will run on (adjust this if needed)
EXPOSE 8080

# Run the application
CMD ["./taurus", "serveHttp"]
