package handler

import (
	"strconv"

	"blog.kimimao.com/backend/models"
	"blog.kimimao.com/backend/service"
	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler 创建
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Register 注册
func (h *UserHandler) Register(c *gin.Context) {
	var req models.UserRegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error400(c, "参数错误")
		return
	}
	u, err := h.userService.Register(&req)
	if err != nil {
		Error(c, 1, err.Error())
		return
	}
	Success(c, u)
}

// Login 登录
func (h *UserHandler) Login(c *gin.Context) {
	var req models.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error400(c, "参数错误")
		return
	}
	token, u, err := h.userService.Login(&req)
	if err != nil {
		Error(c, 1, err.Error())
		return
	}
	Success(c, gin.H{"token": token, "user": u})
}

// GetProfile 获取个人信息
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetInt64("user_id")
	u, err := h.userService.GetProfile(userID)
	if err != nil {
		Error(c, 1, err.Error())
		return
	}
	Success(c, u)
}

// UpdateProfile 更新个人信息
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetInt64("user_id")
	var req models.UserUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error400(c, "参数错误")
		return
	}
	if err := h.userService.UpdateProfile(userID, &req); err != nil {
		Error(c, 1, err.Error())
		return
	}
	SuccessMsg(c, "更新成功")
}

// ChangePassword 修改密码
func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID := c.GetInt64("user_id")
	var req models.ChangePasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error400(c, "参数错误")
		return
	}
	if err := h.userService.ChangePassword(userID, &req); err != nil {
		Error(c, 1, err.Error())
		return
	}
	SuccessMsg(c, "密码修改成功")
}

// List 用户列表（admin）
func (h *UserHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	list, total, err := h.userService.List(page, size)
	if err != nil {
		Error(c, 1, err.Error())
		return
	}
	Success(c, models.PageResult{List: list, Total: total, Page: page, Size: size})
}

// GetUserStats 用户统计
func (h *UserHandler) GetUserStats(c *gin.Context) {
	stats, err := h.userService.GetUserStats()
	if err != nil {
		Error(c, 1, err.Error())
		return
	}
	Success(c, stats)
}
