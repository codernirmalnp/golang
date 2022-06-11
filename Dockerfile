FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
EXPOSE 5000
CMD [ "/app/main" ]
