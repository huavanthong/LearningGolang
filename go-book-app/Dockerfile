# Our builder image used to build the Go binary
FROM golang:1.15.0-alpine as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go mod download

RUN apk add --update --no-cache gcc libc-dev

RUN apk add --update gcc go git mercurial

RUN CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Our production image used to run our app
FROM alpine:latest
# RUN apk update && apk add go gcc bash musl-dev openssl-dev ca-certificates && update-ca-certificates

# Method 1:
# RUN apk --no-cache add ca-certificates
# RUN apk add --no-cache git make musl-dev go

# Method 2:


# RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init 
COPY --from=builder /app/main .

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

# Install Glide
RUN go get -u github.com/Masterminds/glide/...

WORKDIR $GOPATH

CMD ["make"]