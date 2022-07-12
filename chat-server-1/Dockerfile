FROM golang:1.17 as builder

LABEL maintainer="dh.duyhai@gmail.com"

LABEL vendor="MAC DUY HAI"

RUN mkdir /app

WORKDIR /app

COPY . /app



RUN  GO111MODULE=on
RUN  CGO_ENABLED=0
RUN  GOOS=linux
RUN  GOARCH=amd64
RUN  GOFLAGS=-mod=vendor

# RUN ls

# RUN go mod download
RUN go build -o main ./main.go
# RUN go run main.go


FROM ubuntu:16.04

RUN rm -rf /var/lib/apt/lists/*
RUN rm -rf /etc/apt/sources.list.d/*
RUN apt-get update

# RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/*

WORKDIR /app

RUN mkdir storage
RUN mkdir chat_history


COPY --from=builder /app/main .

COPY --from=builder /app/.env .

EXPOSE 80

CMD ["/app/main"]
