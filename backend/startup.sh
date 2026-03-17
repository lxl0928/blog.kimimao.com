#!/bin/bash
# 后端启动脚本

set -e
cd "$(dirname "$0")"

# 检查Go环境
if ! command -v go &> /dev/null; then
    echo "错误: 未安装Go，请先安装Go 1.21+"
    exit 1
fi

# 使用国内代理加速，避免 proxy.golang.org 超时
export GOPROXY=${GOPROXY:-https://goproxy.cn,direct}

# 下载依赖
echo "下载依赖..."
go mod download
go mod tidy

# 可选：运行数据库迁移（需确保PostgreSQL已启动）
# psql -h localhost -U postgres -d blog -f scripts/init.sql 2>/dev/null || true

# 启动服务
echo "启动后端服务..."
go run main.go
