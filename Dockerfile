FROM golang:1.22-alpine

WORKDIR /app

COPY . ./
RUN go mod download


RUN go install github.com/cosmtrek/air@latest


CMD ["air", "-c", ".air.toml"]