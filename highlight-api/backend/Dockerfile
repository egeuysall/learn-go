# Step 1: Use a Go base image to build the app
FROM golang:1.19-alpine as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Download dependencies and build the app
RUN go mod tidy
RUN go build -o highlight-api .

# Step 2: Set up the runtime environment (for the final container)
FROM alpine:latest

# Install CA certificates (needed to access HTTPS resources)
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder container
COPY --from=builder /app/highlight-api /usr/local/bin/highlight-api

# Expose the app's port
EXPOSE 8080

# Start the app when the container is run
CMD ["highlight-api"]
