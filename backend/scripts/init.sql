-- 用户表
CREATE TABLE IF NOT EXISTS tb_user (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    real_name VARCHAR(100),
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(100),
    intro TEXT,
    avatar_url VARCHAR(500),
    role VARCHAR(20) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login_at TIMESTAMP
);

-- 博客表
CREATE TABLE IF NOT EXISTS tb_blog (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES tb_user(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    summary TEXT,
    tags VARCHAR(500),
    view_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 博客评论表
CREATE TABLE IF NOT EXISTS tb_blog_comment (
    id BIGSERIAL PRIMARY KEY,
    blog_id BIGINT NOT NULL REFERENCES tb_blog(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES tb_user(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 博客点赞表
CREATE TABLE IF NOT EXISTS tb_blog_like (
    id BIGSERIAL PRIMARY KEY,
    blog_id BIGINT NOT NULL REFERENCES tb_blog(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES tb_user(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(blog_id, user_id)
);

-- 博客收藏表
CREATE TABLE IF NOT EXISTS tb_blog_favorite (
    id BIGSERIAL PRIMARY KEY,
    blog_id BIGINT NOT NULL REFERENCES tb_blog(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES tb_user(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(blog_id, user_id)
);

-- 索引
CREATE INDEX idx_blog_user_id ON tb_blog(user_id);
CREATE INDEX idx_blog_created_at ON tb_blog(created_at);
CREATE INDEX idx_comment_blog_id ON tb_blog_comment(blog_id);
CREATE INDEX idx_user_created_at ON tb_user(created_at);
CREATE INDEX idx_user_last_login ON tb_user(last_login_at);
