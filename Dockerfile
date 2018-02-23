FROM golang:1.9 as builder

RUN curl https://glide.sh/get | sh

WORKDIR  /go/src/github.com/jnummelin/http-redirector

# Add dependency graph and vendor it in
ADD glide.yaml glide.lock /go/src/github.com/jnummelin/http-redirector/
RUN glide install

# Add source and compile
ADD redirector.go /go/src/github.com/jnummelin/http-redirector/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o redirector .

FROM scratch

COPY --from=builder /go/src/github.com/jnummelin/http-redirector/redirector .

CMD ["./redirector"]