# Blog 后端服务

基于 Go + Gin 的博客 API 服务，提供用户、文章、评论等 RESTful 接口。

## 技术栈

- **框架**: Gin Web Framework
- **数据库**: PostgreSQL
- **缓存**: Redis（登录状态、Token）
- **认证**: JWT

## 目录结构

```
backend/
├── config/          # 配置模块
│   └── config.go    # 应用配置，支持环境变量
├── handler/         # HTTP 处理器层
│   ├── base.go      # 统一响应封装
│   ├── user_handler.go
│   ├── blog_handler.go
│   └── comment_handler.go
├── middleware/      # 中间件
│   ├── auth.go      # JWT 认证、Admin 权限
│   └── cors.go      # 跨域
├── models/          # 数据模型
│   ├── user.go
│   ├── blog.go
│   ├── comment.go
│   └── response.go
├── pkg/             # 公共包
│   ├── database/    # 数据库连接
│   ├── redis/       # Redis 客户端
│   ├── jwt/         # JWT 工具
│   └── password/    # 密码加密
├── repository/      # 数据访问层
│   ├── user_repo.go
│   ├── blog_repo.go
│   ├── comment_repo.go
│   └── like_favorite_repo.go
├── router/          # 路由配置
│   └── router.go
├── scripts/         # 脚本
│   └── init.sql     # 数据库初始化
├── service/         # 业务逻辑层
│   ├── user_service.go
│   ├── blog_service.go
│   └── comment_service.go
├── main.go
├── go.mod
├── startup.sh       # 启动脚本
└── README.md
```

## 依赖

- Go 1.21+
- PostgreSQL 14+
- Redis 6+（可选，连接失败不影响运行）

> 若 `go mod download` 访问 proxy.golang.org 超时，可设置国内代理：
> `export GOPROXY=https://goproxy.cn,direct`

## 环境变量

| 变量 | 说明 | 默认值 |
|------|------|--------|
| SERVER_PORT | 服务端口 | 8080 |
| DB_HOST | 数据库主机 | localhost |
| DB_PORT | 数据库端口 | 5432 |
| DB_USER | 数据库用户 | postgres |
| DB_PASSWORD | 数据库密码 | postgres |
| DB_NAME | 数据库名 | blog |
| REDIS_ADDR | Redis 地址 | localhost:6379 |
| JWT_SECRET | JWT 密钥 | blog-kimimao-secret-key |

## 数据库初始化

```bash
# 创建数据库
createdb blog

# 执行初始化脚本
psql -h localhost -U postgres -d blog -f scripts/init.sql
```

## 启动方式

### 方式一：使用 startup.sh

```bash
chmod +x startup.sh
./startup.sh
```

### 方式二：直接运行

```bash
go mod download
go run main.go
```

### 方式三：编译后运行

```bash
go build -o blog-server .
./blog-server
```

## API 接口概览

### 公开接口
- `POST /api/auth/register` - 注册
- `POST /api/auth/login` - 登录
- `GET /api/blogs` - 文章列表
- `GET /api/blogs/:id` - 文章详情
- `GET /api/blogs/:id/comments` - 评论列表

### 需登录接口
- `GET/PUT /api/user/profile` - 个人信息
- `PUT /api/user/password` - 修改密码
- `POST/GET/PUT/DELETE /api/blogs` - 文章 CRUD
- `POST /api/blogs/:id/comments` - 发表评论
- `POST/DELETE /api/blogs/:id/like` - 点赞
- `POST/DELETE /api/blogs/:id/favorite` - 收藏

### Admin 接口
- `GET /api/admin/users` - 用户列表
- `GET /api/admin/users/stats` - 用户统计
- `GET /api/admin/blogs` - 文章列表
- `GET /api/admin/blogs/stats` - 文章统计
