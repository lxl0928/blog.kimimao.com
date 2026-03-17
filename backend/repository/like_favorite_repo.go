package repository

import (
	"database/sql"
)

// LikeFavoriteRepository 点赞收藏数据访问
type LikeFavoriteRepository struct {
	db *sql.DB
}

// NewLikeFavoriteRepository 创建
func NewLikeFavoriteRepository(db *sql.DB) *LikeFavoriteRepository {
	return &LikeFavoriteRepository{db: db}
}

// Like 点赞
func (r *LikeFavoriteRepository) Like(blogID, userID int64) error {
	_, err := r.db.Exec(`INSERT INTO tb_blog_like (blog_id, user_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`, blogID, userID)
	return err
}

// Unlike 取消点赞
func (r *LikeFavoriteRepository) Unlike(blogID, userID int64) error {
	_, err := r.db.Exec(`DELETE FROM tb_blog_like WHERE blog_id=$1 AND user_id=$2`, blogID, userID)
	return err
}

// IsLiked 是否已点赞
func (r *LikeFavoriteRepository) IsLiked(blogID, userID int64) (bool, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM tb_blog_like WHERE blog_id=$1 AND user_id=$2`, blogID, userID).Scan(&count)
	return count > 0, err
}

// LikeCount 点赞数
func (r *LikeFavoriteRepository) LikeCount(blogID int64) (int, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM tb_blog_like WHERE blog_id=$1`, blogID).Scan(&count)
	return count, err
}

// Favorite 收藏
func (r *LikeFavoriteRepository) Favorite(blogID, userID int64) error {
	_, err := r.db.Exec(`INSERT INTO tb_blog_favorite (blog_id, user_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`, blogID, userID)
	return err
}

// Unfavorite 取消收藏
func (r *LikeFavoriteRepository) Unfavorite(blogID, userID int64) error {
	_, err := r.db.Exec(`DELETE FROM tb_blog_favorite WHERE blog_id=$1 AND user_id=$2`, blogID, userID)
	return err
}

// IsFavorited 是否已收藏
func (r *LikeFavoriteRepository) IsFavorited(blogID, userID int64) (bool, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM tb_blog_favorite WHERE blog_id=$1 AND user_id=$2`, blogID, userID).Scan(&count)
	return count > 0, err
}
