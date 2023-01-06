FROM golang:1.19.4-alpine3.16 as builder

WORKDIR /buildin
COPY . .

RUN go mod download
RUN go build -o /gitstatus main.go

#-------
FROM alpine:latest

WORKDIR /

COPY --from=builder /gitstatus /gitstatus

USER nonroot:nonroot

ENTRYPOINT ["/gitstatus"]