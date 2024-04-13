# Go derleyici image'ını kullanarak başla
FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o kafka_service cmd/main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/kafka_service .

EXPOSE 8081

CMD ["./kafka_service"]
