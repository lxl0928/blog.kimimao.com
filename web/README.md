# Blog 前端

基于 Vue 3 + Vite 的博客前端应用，支持用户端与管理端。

## 技术栈

- **框架**: Vue 3 (Composition API)
- **构建**: Vite 5
- **路由**: Vue Router 4
- **状态**: Pinia
- **HTTP**: Axios
- **Markdown**: marked

## 目录结构

```
web/
├── src/
│   ├── api/              # API 接口模块
│   │   ├── auth.js       # 登录注册
│   │   ├── user.js       # 用户相关
│   │   ├── blog.js       # 文章相关
│   │   ├── comment.js    # 评论相关
│   │   └── admin.js      # 管理端接口
│   ├── components/       # 通用组件
│   │   └── Layout.vue    # 主布局（导航栏）
│   ├── router/           # 路由配置
│   │   └── index.js
│   ├── stores/           # Pinia 状态
│   │   └── user.js       # 用户状态
│   ├── utils/            # 工具
│   │   └── request.js    # Axios 封装
│   ├── views/            # 页面视图
│   │   ├── user/         # 用户端
│   │   │   ├── Login.vue
│   │   │   ├── Register.vue
│   │   │   ├── BlogList.vue
│   │   │   ├── BlogDetail.vue
│   │   │   ├── BlogEdit.vue
│   │   │   └── Profile.vue
│   │   └── admin/        # 管理端
│   │       ├── AdminLayout.vue
│   │       ├── UserList.vue
│   │       └── BlogList.vue
│   ├── App.vue
│   └── main.js
├── index.html
├── package.json
├── vite.config.js
├── startup.sh
└── README.md
```

## 依赖

- Node.js 18+
- npm 或 pnpm

## 启动方式

### 方式一：使用 startup.sh

```bash
chmod +x startup.sh
./startup.sh
```

### 方式二：直接运行

```bash
npm install
npm run dev
```

开发服务默认运行在 http://localhost:3000，API 请求通过 Vite 代理转发到后端 (http://localhost:8080)。

### 生产构建

```bash
npm run build
npm run preview  # 预览构建结果
```

## 功能模块

### 用户端
- 注册、登录、登出
- 修改密码、个人信息
- 文章列表、详情、新增、编辑、删除
- 评论、点赞、收藏、分享链接

### 管理端（需 admin 角色）
- 用户列表及统计（总数、今日新增、今日活跃）
- 文章列表及统计（总数、今日新增、阅读排行）
