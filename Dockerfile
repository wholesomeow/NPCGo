# Use Go 1.24 bookworm as base image
FROM golang:1.24-bookworm AS base

# Development Stage
# -----------------------------------------------------------------------------
# Create development image from base image
FROM base AS development

# Move to working directory /app
WORKDIR /app

# Install air CLI for auto-reloading
RUN go install github.com/air-verse/air@latest

# Copy the go.mod and go.sum files to the /app directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download && go mod verify

# Document the port that may need to be published
EXPOSE 8080

# Start air
CMD [ "air" ]

# Build Stage
# -----------------------------------------------------------------------------
# Create build stage image from base image
FROM base AS build

# Move working dir to /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download && go mod verify

# Copy the entire source code into the container
COPY . .

# Build the application
# Turn off CGO to ensure static binaries
RUN CGO_ENABLED=0 go build -o npcg .

# Production Stage
# -----------------------------------------------------------------------------
# Create a production image to run the application binary
FROM scratch AS production

# Move working dir to /prod
WORKDIR /prod

# Copy binary from build stage
COPY --from=build /build/npcg ./

# Document the port that may need to be published
EXPOSE 8080

# Start the application
CMD ["/prod/npcg"]