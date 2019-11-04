# The following dockerfile is used for producing
# production build artifacts.

# Use the alpine golang image to download dependencies and build the application
FROM golang:alpine as build-env

# Set the Current Working Directory inside the container
ENV PKG_NAME=github.com/gyanesh-mishra/slackbot-winston
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy over source files
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o slackbot-winston cmd/main.go

# Use the base alpine image for the production application image
FROM alpine

# Set the Current Working Directory inside the container
ENV SRC_PATH=/go/src/github.com/gyanesh-mishra/slackbot-winston
ENV DEST_PATH=/opt/slackbot-winston
WORKDIR $DEST_PATH

# Copy over artifacts from builder image
COPY --from=build-env $SRC_PATH $DEST_PATH

ENTRYPOINT ./slackbot-winston