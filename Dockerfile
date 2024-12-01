FROM golang:1.23.3 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .

RUN go build -o /httpserver ./cmd/httpserver/

##
## Deploy
##

FROM ubuntu:22.04

# Install build tools
RUN apt-get update && apt-get install -y wget build-essential

WORKDIR /

COPY --from=build /httpserver /httpserver

EXPOSE 8080

COPY run_server.sh .

RUN chmod +x run_server.sh

CMD ["/bin/bash", "-c", "./run_server.sh"]