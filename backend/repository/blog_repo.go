package repository

import (
	"database/sql"

	"blog.kimimao.com/backend/models"
)

// BlogRepository 博客数据访问
type BlogRepository struct {
	db *sql.DB
}

// NewBlogRepository 创建博客仓库
func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

// Create 创建博客
func (r *BlogRepository) Create(b *models.Blog) error {
	return r.db.QueryRow(
		`INSERT INTO tb_blog (user_id, title, content, summary, tags)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`,
		b.UserID, b.Title, b.Content, b.Summary, b.Tags,
	).Scan(&b.ID, &b.CreatedAt, &b.UpdatedAt)
}

// GetByID 根据ID获取
func (r *BlogRepository) GetByID(id int64) (*models.Blog, error) {
	b := &models.Blog{}
	err := r.db.QueryRow(
		`SELECT id, user_id, title, content, summary, tags, view_count, created_at, updated_at
		 FROM tb_blog WHERE id = $1`, id,
	).Scan(&b.ID, &b.UserID, &b.Title, &b.Content, &b.Summary, &b.Tags, &b.ViewCount, &b.CreatedAt, &b.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return b, err
}

// GetByIDWithAuthor 获取带作者信息
func (r *BlogRepository) GetByIDWithAuthor(id int64) (*models.BlogWithAuthor, error) {
	b := &models.BlogWithAuthor{}
	err := r.db.QueryRow(
		`SELECT b.id, b.user_id, b.title, b.content, b.summary, b.tags, b.view_count, b.created_at, b.updated_at, u.username
		 FROM tb_blog b JOIN tb_user u ON b.user_id = u.id WHERE b.id = $1`, id,
	).Scan(&b.ID, &b.UserID, &b.Title, &b.Content, &b.Summary, &b.Tags, &b.ViewCount, &b.CreatedAt, &b.UpdatedAt, &b.AuthorName)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return b, err
}

// Update 更新博客
func (r *BlogRepository) Update(b *models.Blog) error {
	_, err := r.db.Exec(
		`UPDATE tb_blog SET title=$1, content=$2, summary=$3, tags=$4, updated_at=CURRENT_TIMESTAMP WHERE id=$5 AND user_id=$6`,
		b.Title, b.Content, b.Summary, b.Tags, b.ID, b.UserID,
	)
	return err
}

// Delete 删除博客
func (r *BlogRepository) Delete(id, userID int64) error {
	res, err := r.db.Exec(`DELETE FROM tb_blog WHERE id=$1 AND user_id=$2`, id, userID)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// ListByUser 按用户分页列表（资源归属隔离）
func (r *BlogRepository) ListByUser(userID int64, page, size int) ([]*models.BlogWithAuthor, int64, error) {
	var total int64
	r.db.QueryRow(`SELECT COUNT(*) FROM tb_blog WHERE user_id=$1`, userID).Scan(&total)

	offset := (page - 1) * size
	rows, err := r.db.Query(
		`SELECT b.id, b.user_id, b.title, b.content, b.summary, b.tags, b.view_count, b.created_at, b.updated_at, u.username
		 FROM tb_blog b JOIN tb_user u ON b.user_id = u.id WHERE b.user_id=$1 ORDER BY b.id DESC LIMIT $2 OFFSET $3`,
		userID, size, offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	return scanBlogWithAuthor(rows, total)
}

// ListAll 全部博客分页（公开列表）
func (r *BlogRepository) ListAll(page, size int) ([]*models.BlogWithAuthor, int64, error) {
	var total int64
	r.db.QueryRow(`SELECT COUNT(*) FROM tb_blog`).Scan(&total)

	offset := (page - 1) * size
	rows, err := r.db.Query(
		`SELECT b.id, b.user_id, b.title, b.content, b.summary, b.tags, b.view_count, b.created_at, b.updated_at, u.username
		 FROM tb_blog b JOIN tb_user u ON b.user_id = u.id ORDER BY b.id DESC LIMIT $1 OFFSET $2`,
		size, offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	return scanBlogWithAuthor(rows, total)
}

// IncrViewCount 增加阅读量
func (r *BlogRepository) IncrViewCount(id int64) error {
	_, err := r.db.Exec(`UPDATE tb_blog SET view_count = view_count + 1 WHERE id = $1`, id)
	return err
}

func scanBlogWithAuthor(rows *sql.Rows, total int64) ([]*models.BlogWithAuthor, int64, error) {
	var list []*models.BlogWithAuthor
	for rows.Next() {
		b := &models.BlogWithAuthor{}
		err := rows.Scan(&b.ID, &b.UserID, &b.Title, &b.Content, &b.Summary, &b.Tags, &b.ViewCount, &b.CreatedAt, &b.UpdatedAt, &b.AuthorName)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, b)
	}
	return list, total, nil
}

// CountTotal 文章总数
func (r *BlogRepository) CountTotal() (int64, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM tb_blog`).Scan(&count)
	return count, err
}

// CountTodayNew 今日新增
func (r *BlogRepository) CountTodayNew() (int64, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM tb_blog WHERE DATE(created_at) = CURRENT_DATE`).Scan(&count)
	return count, err
}

// GetViewRank 阅读排行榜（按阅读量排序）
func (r *BlogRepository) GetViewRank(limit int) ([]*models.BlogWithAuthor, error) {
	rows, err := r.db.Query(
		`SELECT b.id, b.user_id, b.title, b.content, b.summary, b.tags, b.view_count, b.created_at, b.updated_at, u.username
		 FROM tb_blog b JOIN tb_user u ON b.user_id = u.id ORDER BY b.view_count DESC LIMIT $1`, limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*models.BlogWithAuthor
	for rows.Next() {
		b := &models.BlogWithAuthor{}
		err := rows.Scan(&b.ID, &b.UserID, &b.Title, &b.Content, &b.Summary, &b.Tags, &b.ViewCount, &b.CreatedAt, &b.UpdatedAt, &b.AuthorName)
		if err != nil {
			return nil, err
		}
		list = append(list, b)
	}
	return list, nil
}
