# Use the official Golang image as a base
FROM golang:1.23

WORKDIR /app

COPY ./app .

RUN go mod tidy
RUN go build -o main .

CMD ["./main"]

EXPOSE 8080
