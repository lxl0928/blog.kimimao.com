#!/bin/bash
# 前端启动脚本

set -e
cd "$(dirname "$0")"

# 检查 Node 环境
if ! command -v npm &> /dev/null; then
    echo "错误: 未安装 Node.js，请先安装 Node 18+"
    exit 1
fi

# 安装依赖
if [ ! -d "node_modules" ]; then
    echo "安装依赖..."
    npm install
fi

# 启动开发服务
echo "启动前端开发服务..."
npm run dev
