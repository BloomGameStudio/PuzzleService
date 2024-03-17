# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.19-buster as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -v -o server

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /app/server

# NOTE: Configuration/Viper is currently not in use in this service.
# Copy config (Even though it is not in use in this service but will throw an error otherwise as it is expected)
# As the binary is being run from the root dir
# the config has to be also in the root dir 
COPY config /config

# Alternativly another viper config path could be added:
# viper.AddConfigPath("/app/config")
# and the config could be copied to the app dir
# COPY config /app/config


# Run the app on container startup.
CMD ["/app/server"]