FROM golang:alpine

# Install glide for package management
RUN apk update && apk add git
RUN go get -u github.com/Masterminds/glide

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/gyanesh-mishra/slackbot-winston

# Copy glide config file and install dependencies
COPY glide.lock .
COPY glide.yaml .

RUN glide install --strip-vendor

# Copy over source files
COPY cmd/ cmd/

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o slackbot-winston cmd/main.go

# Run the application
CMD ["./slackbot-winston"]