FROM golang:1.19.1 AS build

WORKDIR /app

COPY . /app
RUN go mod download

RUN go build -o /app/server cmd/server/main.go

ARG DROPBOX_TOKEN
ENV DROPBOX_TOKEN=$(DROPBOX_TOKEN)

CMD ["/app/server"]
