FROM golang:1.21-alpine AS builder
WORKDIR /usr/local/src/
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /go/bin/app ./main.go

FROM alpine:latest
COPY --from=builder /go/bin/app /
CMD ["/app"]
