package service

import (
	"errors"

	"blog.kimimao.com/backend/models"
	"blog.kimimao.com/backend/repository"
)

// CommentService 评论服务
type CommentService struct {
	commentRepo *repository.CommentRepository
	blogRepo    *repository.BlogRepository
}

// NewCommentService 创建评论服务
func NewCommentService(commentRepo *repository.CommentRepository, blogRepo *repository.BlogRepository) *CommentService {
	return &CommentService{commentRepo: commentRepo, blogRepo: blogRepo}
}

// Create 创建评论
func (s *CommentService) Create(blogID, userID int64, req *models.CommentCreateReq) (*models.BlogComment, error) {
	blog, err := s.blogRepo.GetByID(blogID)
	if err != nil || blog == nil {
		return nil, errors.New("文章不存在")
	}
	c := &models.BlogComment{BlogID: blogID, UserID: userID, Content: req.Content}
	if err := s.commentRepo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

// ListByBlogID 获取文章评论列表
func (s *CommentService) ListByBlogID(blogID int64) ([]*models.CommentWithUser, error) {
	return s.commentRepo.ListByBlogID(blogID)
}
