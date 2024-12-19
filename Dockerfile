# syntax=docker/dockerfile:1

FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY internal ./internal/
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-example
# Run the tests in the container
RUN go test -v ./...

EXPOSE 8000

CMD [ "/docker-example" ]
