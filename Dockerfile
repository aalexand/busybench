FROM golang:1.8

WORKDIR /go/src/busybench
COPY *.go .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
