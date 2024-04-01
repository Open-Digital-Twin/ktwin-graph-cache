FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY internal ./internal
COPY docs ./docs

RUN go mod download
RUN go build -o /ktwin-graph-store

FROM alpine
WORKDIR /app
COPY --from=build /ktwin-graph-store /

EXPOSE 8080

ENTRYPOINT [ "/ktwin-graph-store" ]