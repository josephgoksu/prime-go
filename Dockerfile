# Use an official Golang runtime as a parent image
FROM golang:latest AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go app
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a minimal base image to reduce size
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/main /app/main

# Expose port 8080 for the container
EXPOSE 8080

# Run the Go app when the container starts
CMD ["/app/main"]
