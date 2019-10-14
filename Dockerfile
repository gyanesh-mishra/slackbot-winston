FROM golang:alpine

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/gyanesh-mishra/slackbot-winston

# Copy deps
COPY vendor/ .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# Copy over source files
COPY cmd/ .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o slackbot-winston

# Run the application
CMD ["./slackbot-winston"]