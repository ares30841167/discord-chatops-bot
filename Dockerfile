# Stage1 - Build the executable of disgobot from the golang files
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /build
COPY ./cmd /build/cmd
COPY ./internal /build/internal
COPY ./go.mod /build/
COPY ./go.sum /build/
RUN go mod download
RUN go build -o disgobot ./cmd/disgoBot/main.go

# Stage2 - Build the release image
FROM alpine:latest
WORKDIR /disgobot/
COPY ./assets /disgobot/assets
COPY --from=builder /build/disgobot /disgobot/

CMD ["./disgobot"]