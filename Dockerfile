FROM golang:1.18-rc-bullseye as builder
RUN mkdir -p /app
COPY . /app
WORKDIR /app
RUN go build -o main .

FROM debian:bullseye
RUN mkdir -p /app
COPY --from=builder /app/main /app

WORKDIR /app
ENTRYPOINT  ["/app/main"]
