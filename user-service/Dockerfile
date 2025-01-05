FROM golang:alpine3.19 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o user-service ./user-service/.

FROM alpine:3.19
WORKDIR /user-service/
COPY --from=build app/user-service .
EXPOSE 51152
CMD ["./user-service"]