# 项目主 Dockerfile - 构建后端服务
# 完整编排请使用 docker-compose.yaml
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY backend/go.mod backend/go.sum* ./backend/
WORKDIR /app/backend
RUN go mod download
COPY backend/ .
RUN CGO_ENABLED=0 go build -o server main.go

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/backend/server .
COPY backend/scripts/init.sql ./scripts/
EXPOSE 8080
CMD ["./server"]
