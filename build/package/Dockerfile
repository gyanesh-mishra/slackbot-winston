FROM golang:alpine

# Install glide for package management
RUN apk update && apk add git
RUN go get -u github.com/Masterminds/glide

# Set the Current Working Directory inside the container
ENV PKG_NAME=github.com/gyanesh-mishra/slackbot-winston
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH

# Copy glide config file and install dependencies
COPY glide.lock glide.yaml $PKG_PATH/
RUN glide install --strip-vendor

# Copy over source files
COPY cmd/ .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o slackbot-winston

# Run the application
CMD ["./slackbot-winston"]