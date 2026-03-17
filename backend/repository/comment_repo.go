package repository

import (
	"database/sql"

	"blog.kimimao.com/backend/models"
)

// CommentRepository 评论数据访问
type CommentRepository struct {
	db *sql.DB
}

// NewCommentRepository 创建评论仓库
func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

// Create 创建评论
func (r *CommentRepository) Create(c *models.BlogComment) error {
	return r.db.QueryRow(
		`INSERT INTO tb_blog_comment (blog_id, user_id, content) VALUES ($1, $2, $3) RETURNING id, created_at`,
		c.BlogID, c.UserID, c.Content,
	).Scan(&c.ID, &c.CreatedAt)
}

// ListByBlogID 获取文章评论列表
func (r *CommentRepository) ListByBlogID(blogID int64) ([]*models.CommentWithUser, error) {
	rows, err := r.db.Query(
		`SELECT c.id, c.blog_id, c.user_id, c.content, c.created_at, u.username
		 FROM tb_blog_comment c JOIN tb_user u ON c.user_id = u.id WHERE c.blog_id = $1 ORDER BY c.created_at DESC`,
		blogID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*models.CommentWithUser
	for rows.Next() {
		c := &models.CommentWithUser{}
		err := rows.Scan(&c.ID, &c.BlogID, &c.UserID, &c.Content, &c.CreatedAt, &c.Username)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}
