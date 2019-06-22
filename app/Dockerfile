FROM golang:latest AS builder
ADD . /go/src
WORKDIR /go/src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /server ./
RUN chmod +x ./server
EXPOSE 8080
ENTRYPOINT ["./server"]
