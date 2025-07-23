
FROM golang:latest AS builder

WORKDIR /workspace


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o cloudrun ./cmd


FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /workspace/cloudrun .
COPY --from=builder /workspace/.env .

EXPOSE 8080
ENTRYPOINT ["./cloudrun"]
