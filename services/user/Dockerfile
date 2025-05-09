FROM golang:1.24-bookworm AS build

WORKDIR /app

COPY ./pkg/ /pkg
COPY ./services/user/go.mod ./services/user/go.sum ./
RUN go mod download

COPY ./services/user .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main ./cmd/user

FROM scratch AS final

WORKDIR /

COPY ./services/user/config/config.yaml ./config.yaml
COPY --from=build /bin/main /main

EXPOSE 8080

ENTRYPOINT ["/main"]