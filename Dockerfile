# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
RUN go mod init helm-bug-reproduction
COPY reproduction.go .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o reproduce reproduction.go
FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /app/reproduce .
USER nonroot:nonroot

ENTRYPOINT ["./reproduce"]