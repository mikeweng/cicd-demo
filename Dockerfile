FROM golang:1.13 as builder

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o /go/bin/demo

FROM alpine:3
COPY --from=builder /go/bin/demo /go/bin/demo
CMD ["/go/bin/demo"]
