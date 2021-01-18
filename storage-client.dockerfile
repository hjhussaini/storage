FROM grpc-go:1.15.6-alpine AS build

RUN apk update && apk add build-base

RUN go get github.com/gorilla/mux
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/sqlite
RUN go get github.com/go-playground/validator
RUN go get github.com/dropbox/dropbox-sdk-go-unofficial/dropbox
RUN go get github.com/mitchellh/ioprogress

WORKDIR /go/src/github.com/hjhussaini/storage

COPY protobuf/ .
COPY client/ .

RUN mkdir -p proto && protoc --proto_path=. *.proto --go_out=plugins=grpc:proto
RUN go build -o storage-client

FROM alpine AS runtime

WORKDIR /app

COPY --from=build /go/src/github.com/hjhussaini/storage/storage-client /app

ENTRYPOINT [ "/app/storage-client" ]
