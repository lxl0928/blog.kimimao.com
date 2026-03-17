package models

import "time"

// BlogComment 博客评论模型
type BlogComment struct {
	ID        int64     `json:"id" db:"id"`
	BlogID    int64     `json:"blog_id" db:"blog_id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// CommentWithUser 带用户信息的评论
type CommentWithUser struct {
	BlogComment
	Username string `json:"username"`
}

// CommentCreateReq 创建评论请求
type CommentCreateReq struct {
	Content string `json:"content" binding:"required"`
}
