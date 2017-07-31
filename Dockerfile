FROM golang:1.8

WORKDIR /go/src/busybench
COPY *.go .

EXPOSE 8080

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
