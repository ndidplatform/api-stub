# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/ndid/restapi

WORKDIR "/go/src/github.com/ndid/restapi"

# Build the golang-docker command inside the container.
RUN go install github.com/ndid/restapi

# Run the golang-docker command by default when the container starts.
ENTRYPOINT /go/bin/restapi

# http server listens on port 8000.
EXPOSE 8000
