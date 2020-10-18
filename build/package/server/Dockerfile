FROM golang:1.15-alpine3.12 as base
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go
CMD ["/app/main"]