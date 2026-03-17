package handler

import (
	"strconv"

	"blog.kimimao.com/backend/models"
	"blog.kimimao.com/backend/service"
	"github.com/gin-gonic/gin"
)

// CommentHandler 评论处理器
type CommentHandler struct {
	commentService *service.CommentService
}

// NewCommentHandler 创建
func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

// Create 创建评论
func (h *CommentHandler) Create(c *gin.Context) {
	userID := c.GetInt64("user_id")
	blogID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if blogID <= 0 {
		Error400(c, "无效文章ID")
		return
	}
	var req models.CommentCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error400(c, "参数错误")
		return
	}
	comment, err := h.commentService.Create(blogID, userID, &req)
	if err != nil {
		Error(c, 1, err.Error())
		return
	}
	Success(c, comment)
}

// ListByBlogID 文章评论列表
func (h *CommentHandler) ListByBlogID(c *gin.Context) {
	blogID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if blogID <= 0 {
		Error400(c, "无效文章ID")
		return
	}
	list, err := h.commentService.ListByBlogID(blogID)
	if err != nil {
		Error(c, 1, err.Error())
		return
	}
	Success(c, list)
}
