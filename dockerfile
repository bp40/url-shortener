FROM golang:1.23.0-alpine

WORKDIR /app
COPY . .

ENV CGO_ENABLED=1

RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev

RUN go mod download
RUN go build -o /app/url-shortener

EXPOSE 8080

RUN go build -o bin .

ENTRYPOINT ["/app/bin"]