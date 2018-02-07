FROM golang:1.9 as builder

WORKDIR  /go/src/github.com/jnummelin/go-inception-client

ADD redirector.go /go/src/github.com/jnummelin/go-inception-client/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o redirector .



FROM scratch

COPY --from=builder /go/src/github.com/jnummelin/go-inception-client/redirector .

CMD ["./redirector"]