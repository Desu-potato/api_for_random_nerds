# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /


COPY go.mod .
COPY go.sum .
RUN go mod download
COPY main.go .
RUN mkdir helpers
COPY ./helpers/* ./helpers/
RUN go build -o /api_for_random_nerds
ENV API_KEY_RANDOM ""
ENTRYPOINT ["./api_for_random_nerds"]
