# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY . .
RUN go mod download

WORKDIR /app/cmd/app
RUN GOOS=linux go build -o app

EXPOSE 8080

CMD ["./app"]

# docker build -f .\build\Dockerfile --progress=plain -t dmitriytitov/microapp:latest .
# docker run --rm -p 8080:8080 dmitriytitov/microapp:latest