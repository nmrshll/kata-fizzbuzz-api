FROM golang:1.11.1-alpine3.7 AS build
RUN apk --no-cache add gcc g++ git libc-dev make ca-certificates
WORKDIR /go/src/gitlab.com/nmrshll/go-api-postgres

ENV GO111MODULE=on

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . .
RUN go install ./... 

#

FROM alpine:3.7 AS run-base
RUN apk add bash curl ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/bin .