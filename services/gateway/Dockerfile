FROM golang:1.24-bookworm AS build

WORKDIR /app

COPY ./pkg/ /pkg
COPY ./services/gateway/go.mod ./services/gateway/go.sum ./
RUN go mod download

COPY ./services/gateway .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main ./cmd/gateway

FROM scratch AS final

WORKDIR /

COPY ./services/gateway/config/config.yaml ./config.yaml
COPY --from=build /bin/main /main

EXPOSE 8080

ENTRYPOINT ["/main"]