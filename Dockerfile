FROM golang:alpine

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o brun

WORKDIR /dist

RUN cp /build/brun .

EXPOSE 8080

CMD ["/dist/brun"]