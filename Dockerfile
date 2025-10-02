# Frontend builder
FROM oven/bun:1.2-alpine AS frontend-builder
WORKDIR /app/frontend
COPY apps/frontend/package.json ./
RUN bun ci
COPY apps/frontend/ ./
RUN bun run build

# Backend builder
FROM golang:1.25-alpine AS backend-builder
WORKDIR /app/backend
COPY apps/backend/go.mod ./
COPY apps/backend/go.sum ./
RUN go mod download

COPY apps/backend/ ./
RUN go build -o api ./cmd/api

# Final image
FROM alpine:latest
WORKDIR /app

# Copy backend
COPY --from=backend-builder /app/backend/api ./main
# Copy frontend build
COPY --from=frontend-builder /app/frontend/build/ ./public/

EXPOSE 3000
CMD ["./main"]