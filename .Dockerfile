FROM golang:1.24 AS builder

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=1 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o app main.go

FROM debian:stable-slim

WORKDIR /server

COPY --from=builder /server/app ./app
COPY templates templates
COPY static static
COPY db/migration ./db/migration

CMD ["./app"]