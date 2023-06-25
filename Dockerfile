FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum /app

RUN go mod download && go mod verify

COPY . .

EXPOSE 8080

CMD ["go", "run", "main.go"]
