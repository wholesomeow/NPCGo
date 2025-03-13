# Use Go 1.23 alpine as base image
FROM golang:1.23-alpine AS base

# Move to working directory /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download && go mod verify

# Copy the entire source code into the container
COPY . .

# Build the application
RUN go build -o npcg .

# Document the port that may need to be published
EXPOSE 8000

# Start the application
CMD ["/build/npcg"]