FROM golang:1.20-alpine AS build

WORKDIR /app

RUN apk --update add ca-certificates upx && update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -ldflags="-s -w" -o /bin/mymm . && upx --best --lzma /bin/mymm

FROM alpine:3.18.0 AS release

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /bin/mymm /app/mymm
COPY config.toml /app/config.toml
COPY users.yml /app/users.yml

ENTRYPOINT ["/app/mymm"]
CMD ["--config", "/app/config.toml"]
