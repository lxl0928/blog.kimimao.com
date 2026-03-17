package handler

import (
	"net/http"

	"blog.kimimao.com/backend/models"
	"github.com/gin-gonic/gin"
)

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.Response{Code: 0, Message: "success", Data: data})
}

// SuccessMsg 成功消息
func SuccessMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, models.Response{Code: 0, Message: msg})
}

// Error 错误响应
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, models.Response{Code: code, Message: msg})
}

// Error400 参数错误
func Error400(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, models.Response{Code: 400, Message: msg})
}
