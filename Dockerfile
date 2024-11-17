FROM golang:1.20-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .

RUN go build -o /httpserver ./cmd/httpserver/

##
## Deploy
##

# FROM gcr.io/distroless/base-debian10
FROM debian:11

#Install Shell
# RUN apt-get install debian-keyring debian-archive-keyring -y \
#     && apt-key update \
#     && apt-get update
RUN apt-get update && apt-get install -y \
    build-essential \
    wget \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /

COPY --from=build /httpserver /httpserver

EXPOSE 8080

COPY run_servers.sh .

RUN chmod +x run_servers.sh

CMD ["/bin/bash", "-c", "./run_servers.sh"]