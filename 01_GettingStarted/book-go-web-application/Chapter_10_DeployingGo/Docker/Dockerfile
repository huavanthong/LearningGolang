# Starts from a Debian image with Go installed and a GOPATH configured at /go.
FROM golang

# Copies the local package files to the container’s workspace.
RUN mkdir /app
ADD . /app
WORKDIR /app

# Builds the ws-dcommand inside the container.
RUN go get github.com/lib/pq
RUN go install server

# Runs the ws-d command by default when the container starts.
ENTRYPOINT /go/bin/app

# Documents that the service listens on port 8080.
EXPOSE 8080