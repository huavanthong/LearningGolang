FROM node:latest AS node_builder

RUN mkdir -p /app/webreactjs
WORKDIR /app/webreactjs
COPY /webreactjs/package.json .

# Install npm 
RUN npm install
# Beacause node:latest does not use apk.
# Install bash for webreactjs container 
# Refer: [here](https://stackoverflow.com/questions/58210383/nodelatest-for-alpine-apk-not-found-because-sbin-is-not-on-path)
RUN apt update && apt add bash

COPY /webreactjs .
LABEL name="webreactjs" version="1.0"
EXPOSE 3000
CMD ["npm", "start"]

FROM golang:1.17 AS go_builder
ADD . /app
WORKDIR /app
# Install bash for servergo container
RUN apk update && apk add bash
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .
LABEL name="server" version="1.0"

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=go_builder /main ./
RUN chmod +x ./main
EXPOSE 8080
CMD ./main
