FROM golang:1.17.5-alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build .

FROM alpine
RUN mkdir /app
WORKDIR /app
RUN mkdir -p schemas/migrations
COPY --from=builder /app/UwU .
COPY --from=builder /app/schemas/migrations/* schemas/migrations/

EXPOSE 3000

CMD ["./UwU"]