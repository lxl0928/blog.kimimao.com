# 个人博客 - Blog.kimimao.com

Go + Vue3 前后端分离的个人博客项目。

## 技术栈

| 层级 | 技术 |
|------|------|
| 后端 | Golang + Gin |
| 数据库 | PostgreSQL |
| 缓存 | Redis（登录状态、Token） |
| 前端 | Vue 3 + Vite |

## 项目结构

```
blog.kimimao.com/
├── backend/          # 后端 Go 服务
│   ├── config/       # 配置
│   ├── handler/      # HTTP 处理器
│   ├── middleware/   # 中间件
│   ├── models/       # 数据模型
│   ├── pkg/          # 公共包
│   ├── repository/   # 数据访问层
│   ├── router/       # 路由
│   ├── service/      # 业务逻辑
│   ├── scripts/      # 初始化脚本
│   ├── main.go
│   ├── startup.sh
│   └── README.md
├── web/              # 前端 Vue 应用
│   ├── src/
│   │   ├── api/      # API 接口
│   │   ├── components/
│   │   ├── views/    # 页面（user/admin）
│   │   ├── router/
│   │   └── stores/
│   ├── startup.sh
│   └── README.md
├── Dockerfile        # 后端镜像构建
├── docker-compose.yaml
└── README.md
```

## 功能模块

### 用户端 (user)
- 注册、登录、登出、修改密码
- 个人信息（用户名、真实姓名、手机、邮箱、简介、头像）
- 文章列表、详情、新增、编辑、删除（按 user_id 权限隔离）
- 评论、点赞、收藏、分享链接

### 管理端 (admin)
- 用户列表：用户总数、今日新增、今日活跃
- 文章列表：文章总数、今日新增、阅读排行

## 依赖

- Go 1.21+
- Node.js 18+
- PostgreSQL 14+
- Redis 6+（可选）

## 启动方式

### 方式一：Docker Compose（推荐）

```bash
docker-compose up -d
```

- 前端: http://localhost
- 后端 API: http://localhost:8080
- PostgreSQL: localhost:5432
- Redis: localhost:6379

### 方式二：本地开发

**1. 启动数据库**
```bash
# PostgreSQL
createdb blog
psql -d blog -f backend/scripts/init.sql

# Redis（可选）
redis-server
```

**2. 启动后端**
```bash
cd backend
./startup.sh
# 或: go run main.go
```

**3. 启动前端**
```bash
cd web
./startup.sh
# 或: npm install && npm run dev
```

前端开发服务: http://localhost:3000（API 代理到 8080）

## 环境变量

| 变量 | 说明 | 默认值 |
|------|------|--------|
| SERVER_PORT | 后端端口 | 8080 |
| DB_HOST | 数据库主机 | localhost |
| DB_PORT | 数据库端口 | 5432 |
| DB_USER | 数据库用户 | postgres |
| DB_PASSWORD | 数据库密码 | postgres |
| DB_NAME | 数据库名 | blog |
| REDIS_ADDR | Redis 地址 | localhost:6379 |
| JWT_SECRET | JWT 密钥 | blog-kimimao-secret-key |

## 创建管理员

数据库初始化后，需手动将用户角色改为 admin：

```sql
UPDATE tb_user SET role = 'admin' WHERE username = 'your_username';
```

## 文档

- [后端文档](backend/README.md)
- [前端文档](web/README.md)
