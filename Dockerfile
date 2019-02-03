FROM golang:1.11 as build
WORKDIR /go/src/github.com/jukeizu/rand
COPY Makefile go.mod go.sum ./
RUN make deps
ADD . .
RUN make build-linux
RUN echo "rand:x:100:101:/" > passwd

FROM scratch
COPY --from=build /go/src/github.com/jukeizu/rand/passwd /etc/passwd
COPY --from=build --chown=100:101 /go/src/github.com/jukeizu/rand/bin/rand .
USER rand
ENTRYPOINT ["./rand"]
