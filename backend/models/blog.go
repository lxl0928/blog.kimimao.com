package models

import "time"

// Blog 博客文章模型
type Blog struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Summary   string    `json:"summary" db:"summary"`
	Tags      string    `json:"tags" db:"tags"`
	ViewCount int       `json:"view_count" db:"view_count"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// BlogWithAuthor 带作者信息的博客
type BlogWithAuthor struct {
	Blog
	AuthorName string `json:"author_name"`
}

// BlogCreateReq 创建文章请求
type BlogCreateReq struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Summary string `json:"summary"`
	Tags    string `json:"tags"`
}

// BlogUpdateReq 更新文章请求
type BlogUpdateReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Summary string `json:"summary"`
	Tags    string `json:"tags"`
}
