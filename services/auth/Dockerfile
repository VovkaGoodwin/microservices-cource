FROM golang:1.24-bookworm AS build

WORKDIR /app

COPY ./pkg/ /pkg
COPY ./services/auth/go.mod ./services/auth/go.sum ./
RUN go mod download

COPY ./services/auth .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main ./cmd/auth

FROM scratch AS final

WORKDIR /

COPY ./services/auth/config/config.yaml ./config.yaml
COPY --from=build /bin/main /main

EXPOSE 8080

ENTRYPOINT ["/main"]