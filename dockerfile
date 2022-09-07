FROM golang:alpine3.15 as builder
RUN mkdir /app
WORKDIR /app
COPY ./* /app/
RUN go build main.go
FROM alpine as application
COPY --from=builder /app/main /app
WORKDIR /app
ENTRYPOINT [ "/app/main" ]
