FROM golang:latest

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . .

RUN go build -o /run-api

EXPOSE 3000

CMD ["/run-api"]