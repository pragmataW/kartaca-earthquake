FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o filtering cmd/main.go

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/filtering .

EXPOSE 5435

CMD ["./filtering"]
