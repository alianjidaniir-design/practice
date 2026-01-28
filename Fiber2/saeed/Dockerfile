FROM golang:1.25.3-alpine

WORKDIR /fib

COPY go.mod go.sum ./
RUN go mod download

COPY  .  .

RUN go build -o f2 ./Fiber2/saeed

EXPOSE 3000

CMD ["./saeed"]
