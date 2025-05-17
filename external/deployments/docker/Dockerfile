# ======================
#  GO STAGE
# ======================
FROM golang:latest AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o main ./cmd/api

# ======================
#  ALPINE STAGE
# ======================
FROM alpine:latest
WORKDIR /usr/src/app

COPY --from=builder /app/main .

EXPOSE 3000
ENTRYPOINT ["./main"]