# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o task main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/task .
COPY app.env .
COPY start.sh .

EXPOSE 8080
CMD [ "/app/task" ]
ENTRYPOINT [ "/app/start.sh" ]
