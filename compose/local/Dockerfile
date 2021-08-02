FROM golang:lastest As builder
RUN apt-get update

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/app
COPY go.mode .
RUN go mod dowload
COPY . .
RUN go install
CMD ["./main"]
