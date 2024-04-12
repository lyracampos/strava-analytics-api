FROM golang:1.22.2 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /strava-analytics-api

FROM alpine

WORKDIR /

COPY --from=build /strava-analytics-api /strava-analytics-api

EXPOSE 8080

ENTRYPOINT [ "/strava-analytics-api" ]
