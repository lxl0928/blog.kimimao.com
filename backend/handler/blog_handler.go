package handler

import (
	"strconv"

	"blog.kimimao.com/backend/models"
	"blog.kimimao.com/backend/service"
	"github.com/gin-gonic/gin"
)

// BlogHandler 博客处理器
type BlogHandler struct {
	blogService *service.BlogService
}

// NewBlogHandler 创建
func NewBlogHandler(blogService *service.BlogService) *BlogHandler {
	return &BlogHandler{blogService: blogService}
}

// Create 创建文章
func (h *BlogHandler) Create(c *gin.Context) {
	userID := c.GetInt64("user_id")
	var req models.BlogCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error400(c, "参数错误")
		return
	}
	b, err := h.blogService.Create(userID, &req)
	if err != nil {
		Error(c, 1, err.Error())
		return
	}
	Success(c, b)
}

// GetDetail 文章详情
func (h *BlogHandler) GetDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 {
		Error400(c, "无效ID")
		return
	}
	incrView := c.Query("incr_view") != "false"
	b, err := h.blogService.GetDetail(id, incrView)
	if err != nil {
		Error(c, 1, err.Error())
		return
	}
	Success(c, b)
}

// Update 更新文章
func (h *BlogHandler) Update(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 {
		Error400(c, "无效ID")
		return
	}
	var req models.BlogUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error400(c, "参数错误")
		return
	}
	if err := h.blogService.Update(userID, id, &req); err != nil {
		Error(c, 1, err.Error())
		return
	}
	SuccessMsg(c, "更新成功")
}

// Delete 删除文章
func (h *BlogHandler) Delete(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 {
		Error400(c, "无效ID")
		return
	}
	if err := h.blogService.Delete(userID, id); err != nil {
		Error(c, 1, err.Error())
		return
	}
	SuccessMsg(c, "删除成功")
}

// ListMy 我的文章列表
func (h *BlogHandler) ListMy(c *gin.Context) {
	userID := c.GetInt64("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	list, total, err := h.blogService.ListByUser(userID, page, size)
	if err != nil {
		Error(c, 1, err.Error())
		return
	}
	Success(c, models.PageResult{List: list, Total: total, Page: page, Size: size})
}

// ListAll 公开文章列表
func (h *BlogHandler) ListAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	list, total, err := h.blogService.ListAll(page, size)
	if err != nil {
		Error(c, 1, err.Error())
		return
	}
	Success(c, models.PageResult{List: list, Total: total, Page: page, Size: size})
}

// Like 点赞
func (h *BlogHandler) Like(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.blogService.Like(id, userID); err != nil {
		Error(c, 1, err.Error())
		return
	}
	SuccessMsg(c, "点赞成功")
}

// Unlike 取消点赞
func (h *BlogHandler) Unlike(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.blogService.Unlike(id, userID); err != nil {
		Error(c, 1, err.Error())
		return
	}
	SuccessMsg(c, "已取消点赞")
}

// Favorite 收藏
func (h *BlogHandler) Favorite(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.blogService.Favorite(id, userID); err != nil {
		Error(c, 1, err.Error())
		return
	}
	SuccessMsg(c, "收藏成功")
}

// Unfavorite 取消收藏
func (h *BlogHandler) Unfavorite(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.blogService.Unfavorite(id, userID); err != nil {
		Error(c, 1, err.Error())
		return
	}
	SuccessMsg(c, "已取消收藏")
}

// GetBlogStats 文章统计（admin）
func (h *BlogHandler) GetBlogStats(c *gin.Context) {
	stats, err := h.blogService.GetBlogStats(10)
	if err != nil {
		Error(c, 1, err.Error())
		return
	}
	Success(c, stats)
}
