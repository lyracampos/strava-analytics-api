FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./strava-analytics-api ./cmd

FROM alpine
COPY --from=build /app/strava-analytics-api /usr/local/bin/app
COPY --from=build /app/config /config

ENTRYPOINT ["app", "-c", "/config/config.yaml"]

EXPOSE 8080