FROM golang:1.14-alpine as builder

WORKDIR /app

COPY go.* ./

ENV GOOS=linux

RUN go mod download & go mod verify

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/nats-streams ./main.go

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=builder /app/nats-streams .

ENTRYPOINT ["/app/nats-streams"]