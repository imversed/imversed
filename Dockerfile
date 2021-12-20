FROM golang:1.16.5-alpine AS build-env

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/cosmos/cosmos-sdk

# Add source files
COPY . .

COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# install simapp, remove packages
RUN make build-linux

RUN ls -a
RUN ls -a build/

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env

COPY --from=build-env /go/src/github.com/cosmos/cosmos-sdk/build/imversed /usr/bin/imversed

EXPOSE 26656 26657 1317 9090

CMD ["imversed", "start"]
