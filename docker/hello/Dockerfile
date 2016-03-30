FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/hello

# Build the golang-docker command inside the container.
RUN go install hello

# Run the hello command when the container starts.
ENTRYPOINT /go/bin/hello

# http server listens on port 7000.
EXPOSE 7000