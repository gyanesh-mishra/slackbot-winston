FROM golang:alpine

# Set the Current Working Directory inside the container
ENV PKG_NAME=github.com/gyanesh-mishra/slackbot-winston
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH

# Copy over source files
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o slackbot-winston cmd/main.go

# Run the application
CMD ["./slackbot-winston"]