FROM golang:1.20-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download && go build -o p100-api

FROM alpine

COPY --from=builder /build/p100-api /app/p100-api

CMD ["/app/p100-api"]
