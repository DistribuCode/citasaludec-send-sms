FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o send-sms ./cmd/main.go && ls -l

EXPOSE 5000

CMD ["./send-sms"]
