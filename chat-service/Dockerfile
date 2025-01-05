FROM golang:alpine3.19 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o chat-service ./chat-service/.

FROM alpine:3.19
WORKDIR /chat-service/
COPY --from=build app/chat-service .
EXPOSE 51153
CMD ["./chat-service"]