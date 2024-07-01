# Dockerfile
# Use the official Golang image to create a build artifact.
FROM golang:1.19-alpine AS build

# Create and change to the app directory.
WORKDIR /app

# Copy Go module files.
COPY go.mod go.sum ./

# Download dependencies.
RUN go mod download

# Copy the source code.
COPY . .

# Build the application.
RUN go build -o main .

# Use a minimal image for the final container.
FROM alpine:latest

# Set the working directory.
WORKDIR /root/

# Copy the build artifact from the build stage.
COPY --from=build /app/main .

# Run the application.
CMD ["./main"]
