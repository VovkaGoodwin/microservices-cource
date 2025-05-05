FROM golang:1.24-bookworm AS build

WORKDIR /app

COPY ./pkg/ /pkg
COPY ./services/message/go.mod ./services/message/go.sum ./
RUN go mod download

COPY ./services/message .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main ./cmd/message

FROM scratch AS final

WORKDIR /

COPY ./services/message/config/config.yaml ./config.yaml
COPY --from=build /bin/main /main

EXPOSE 8080

ENTRYPOINT ["/main"]