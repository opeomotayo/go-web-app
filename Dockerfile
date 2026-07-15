# ---- build stage ----
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o go-web-app main.go

# ---- runtime stage ----
FROM alpine:3.20
RUN adduser -D -g '' appuser
WORKDIR /app
COPY --from=builder /app/go-web-app .
USER appuser
EXPOSE 8080
ENTRYPOINT ["./go-web-app"]
