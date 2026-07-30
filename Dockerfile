# ── Stage 1: Build static Go production binary ──
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Install ca-certificates and git
RUN apk add --no-cache ca-certificates git

# Copy dependency files and download modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Compile static binary optimized for size (-w -s strips debug symbols)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# ── Stage 2: Minimal runtime image ──
FROM alpine:latest

WORKDIR /app

# Install CA certificates for HTTPS APIs (OpenWeatherMap, GitHub, TabNews) and tzdata
RUN apk add --no-cache ca-certificates tzdata

# Create data directory for notes and wallpaper persistence
RUN mkdir -p /app/data

# Copy binary and template/static directories
COPY --from=builder /app/main /app/main
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static

EXPOSE 3000

ENV PORT=3000

ENTRYPOINT ["/app/main"]
