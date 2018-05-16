FROM golang:alpine

WORKDIR /go/src/busybench
COPY *.go .

EXPOSE 8080

RUN apk update \
    && apk add --no-cache git \
    && go get -d ./... \
    && apk del git

RUN go install ./...

CMD ["busybench"]
