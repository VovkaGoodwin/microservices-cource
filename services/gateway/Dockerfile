FROM golang:1.24-bookworm AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main ./cmd/gateway

FROM scratch AS final

WORKDIR /

COPY --from=build /bin/main /main

EXPOSE 8080

ENTRYPOINT ["/main"]