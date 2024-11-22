FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./apps/model

FROM alpine:3.18.2

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/rpc/conf-release.yaml ./rpc/conf-release.yaml
COPY --from=builder /app/apps/model/conf/conf-release.yaml ./apps/model/conf/conf-release.yaml

EXPOSE 80

ARG GO_ENV=release
ENV GO_ENV=$GO_ENV

VOLUME /app/log

CMD ["./main"]
