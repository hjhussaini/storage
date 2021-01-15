FROM grpc-go:1.15.6-alpine AS build

WORKDIR /go/src/github.com/hjhussaini/storage

COPY protobuf/ .
COPY server/ .

RUN mkdir -p proto && protoc --proto_path=. *.proto --go_out=plugins=grpc:proto
RUN go build -o storage-server

FROM alpine AS runtime

WORKDIR /app

COPY --from=build /go/src/github.com/hjhussaini/storage/storage-server /app

ENTRYPOINT [ "/app/storage-server" ]
