FROM golang:1.25.3-alpine

WORKDIR /Fiber2

COPY go.mod go.sum ./
RUN go mod download

COPY  .  .

RUN go build -o saeed .

EXPOSE 3000

CMD ["./saeed"]
