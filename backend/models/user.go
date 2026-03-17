package models

import "time"

// Role 用户角色
const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

// User 用户模型
type User struct {
	ID         int64     `json:"id" db:"id"`
	Username   string    `json:"username" db:"username"`
	RealName   string    `json:"real_name" db:"real_name"`
	Password   string    `json:"-" db:"password"`
	Phone      string    `json:"phone" db:"phone"`
	Email      string    `json:"email" db:"email"`
	Intro      string    `json:"intro" db:"intro"`
	AvatarURL  string    `json:"avatar_url" db:"avatar_url"`
	Role       string    `json:"role" db:"role"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
}

// UserRegisterReq 注册请求
type UserRegisterReq struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// UserLoginReq 登录请求
type UserLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserUpdateReq 用户信息更新请求
type UserUpdateReq struct {
	RealName  string `json:"real_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Intro     string `json:"intro"`
	AvatarURL string `json:"avatar_url"`
}

// ChangePasswordReq 修改密码请求
type ChangePasswordReq struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}
