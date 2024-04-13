FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o earthquake cmd/main.go

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/earthquake .

EXPOSE 8080

CMD ["./earthquake"]
