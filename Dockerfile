FROM golang:1.24.1-alpine

WORKDIR /device-manager

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -a -o ./bin/api ./cmd/api

CMD ["/device-manager/api"]
EXPOSE 8080
