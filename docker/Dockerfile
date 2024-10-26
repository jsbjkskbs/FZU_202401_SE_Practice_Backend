FROM golang:1.23-alpine AS builder

COPY biz/ /go/src/biz/
COPY pkg/ /go/src/pkg/
COPY go.mod go.sum *.go /go/src/

ENV GO111MODULE=on \
    GOOS=linux \
    FOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /go/src/
RUN go mod tidy
RUN go build -o main

FROM jrottenberg/ffmpeg:4.4-alpine

RUN mkdir "/app"
COPY --from=builder /go/src/main /app/main

RUN chmod +x /app/main

ENTRYPOINT [ "/app/main" ]