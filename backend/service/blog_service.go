package service

import (
	"errors"

	"blog.kimimao.com/backend/models"
	"blog.kimimao.com/backend/repository"
)

// BlogService 博客服务
type BlogService struct {
	blogRepo   *repository.BlogRepository
	likeRepo   *repository.LikeFavoriteRepository
}

// NewBlogService 创建博客服务
func NewBlogService(blogRepo *repository.BlogRepository, likeRepo *repository.LikeFavoriteRepository) *BlogService {
	return &BlogService{blogRepo: blogRepo, likeRepo: likeRepo}
}

// Create 创建文章
func (s *BlogService) Create(userID int64, req *models.BlogCreateReq) (*models.Blog, error) {
	b := &models.Blog{
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
		Summary: req.Summary,
		Tags:    req.Tags,
	}
	if err := s.blogRepo.Create(b); err != nil {
		return nil, err
	}
	return b, nil
}

// GetDetail 获取详情（增加阅读量）
func (s *BlogService) GetDetail(id int64, incrView bool) (*models.BlogWithAuthor, error) {
	b, err := s.blogRepo.GetByIDWithAuthor(id)
	if err != nil || b == nil {
		return nil, errors.New("文章不存在")
	}
	if incrView {
		s.blogRepo.IncrViewCount(id)
		b.ViewCount++
	}
	return b, nil
}

// Update 更新文章（权限校验：仅作者可编辑）
func (s *BlogService) Update(userID, id int64, req *models.BlogUpdateReq) error {
	b, err := s.blogRepo.GetByID(id)
	if err != nil || b == nil {
		return errors.New("文章不存在")
	}
	if b.UserID != userID {
		return errors.New("无权限编辑")
	}
	b.Title = req.Title
	b.Content = req.Content
	b.Summary = req.Summary
	b.Tags = req.Tags
	return s.blogRepo.Update(b)
}

// Delete 删除文章
func (s *BlogService) Delete(userID, id int64) error {
	return s.blogRepo.Delete(id, userID)
}

// ListByUser 我的文章列表
func (s *BlogService) ListByUser(userID int64, page, size int) ([]*models.BlogWithAuthor, int64, error) {
	if size <= 0 {
		size = 10
	}
	return s.blogRepo.ListByUser(userID, page, size)
}

// ListAll 公开文章列表
func (s *BlogService) ListAll(page, size int) ([]*models.BlogWithAuthor, int64, error) {
	if size <= 0 {
		size = 10
	}
	return s.blogRepo.ListAll(page, size)
}

// Like 点赞
func (s *BlogService) Like(blogID, userID int64) error {
	return s.likeRepo.Like(blogID, userID)
}

// Unlike 取消点赞
func (s *BlogService) Unlike(blogID, userID int64) error {
	return s.likeRepo.Unlike(blogID, userID)
}

// Favorite 收藏
func (s *BlogService) Favorite(blogID, userID int64) error {
	return s.likeRepo.Favorite(blogID, userID)
}

// Unfavorite 取消收藏
func (s *BlogService) Unfavorite(blogID, userID int64) error {
	return s.likeRepo.Unfavorite(blogID, userID)
}

// GetBlogStats 文章统计（admin）
func (s *BlogService) GetBlogStats(rankLimit int) (*models.StatsResult, error) {
	total, _ := s.blogRepo.CountTotal()
	todayNew, _ := s.blogRepo.CountTodayNew()
	rankList, _ := s.blogRepo.GetViewRank(rankLimit)
	return &models.StatsResult{
		TotalCount:    total,
		TodayNewCount: todayNew,
		RankList:      rankList,
	}, nil
}
