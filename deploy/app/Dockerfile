FROM golang:1.24.4-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o weather-bot cmd/app/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/weather-bot .
COPY --from=build /app/migrations ./migrations

RUN apk --no-cache add ca-certificates tzdata

ENTRYPOINT ["/app/weather-bot"]