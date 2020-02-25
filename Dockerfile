FROM golang:latest

WORKDIR /app/microservices/

COPY ./ /app/microservices/

RUN go build -o microservices

EXPOSE 8080

CMD ["./microservices"]

