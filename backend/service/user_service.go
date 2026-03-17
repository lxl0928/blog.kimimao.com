package service

import (
	"errors"

	"blog.kimimao.com/backend/config"
	"blog.kimimao.com/backend/models"
	"blog.kimimao.com/backend/pkg/jwt"
	"blog.kimimao.com/backend/pkg/password"
	"blog.kimimao.com/backend/repository"
)

// UserService 用户服务
type UserService struct {
	userRepo *repository.UserRepository
	cfg      *config.Config
}

// NewUserService 创建用户服务
func NewUserService(userRepo *repository.UserRepository, cfg *config.Config) *UserService {
	return &UserService{userRepo: userRepo, cfg: cfg}
}

// Register 注册
func (s *UserService) Register(req *models.UserRegisterReq) (*models.User, error) {
	exist, _ := s.userRepo.GetByUsername(req.Username)
	if exist != nil {
		return nil, errors.New("用户名已存在")
	}
	hash, err := password.Hash(req.Password)
	if err != nil {
		return nil, err
	}
	u := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hash,
		Role:     models.RoleUser,
	}
	if err := s.userRepo.Create(u); err != nil {
		return nil, err
	}
	u.Password = ""
	return u, nil
}

// Login 登录
func (s *UserService) Login(req *models.UserLoginReq) (string, *models.User, error) {
	u, err := s.userRepo.GetByUsername(req.Username)
	if err != nil || u == nil {
		return "", nil, errors.New("用户名或密码错误")
	}
	if !password.Check(req.Password, u.Password) {
		return "", nil, errors.New("用户名或密码错误")
	}
	s.userRepo.UpdateLastLogin(u.ID)
	token, err := jwt.GenerateToken(u.ID, u.Username, u.Role, &s.cfg.JWT)
	if err != nil {
		return "", nil, err
	}
	u.Password = ""
	return token, u, nil
}

// GetProfile 获取个人信息
func (s *UserService) GetProfile(userID int64) (*models.User, error) {
	u, err := s.userRepo.GetByID(userID)
	if err != nil || u == nil {
		return nil, errors.New("用户不存在")
	}
	u.Password = ""
	return u, nil
}

// UpdateProfile 更新个人信息
func (s *UserService) UpdateProfile(userID int64, req *models.UserUpdateReq) error {
	u, err := s.userRepo.GetByID(userID)
	if err != nil || u == nil {
		return errors.New("用户不存在")
	}
	u.RealName = req.RealName
	u.Phone = req.Phone
	u.Email = req.Email
	u.Intro = req.Intro
	u.AvatarURL = req.AvatarURL
	return s.userRepo.Update(u)
}

// ChangePassword 修改密码
func (s *UserService) ChangePassword(userID int64, req *models.ChangePasswordReq) error {
	u, err := s.userRepo.GetByID(userID)
	if err != nil || u == nil {
		return errors.New("用户不存在")
	}
	if !password.Check(req.OldPassword, u.Password) {
		return errors.New("原密码错误")
	}
	hash, err := password.Hash(req.NewPassword)
	if err != nil {
		return err
	}
	return s.userRepo.UpdatePassword(userID, hash)
}

// List 用户列表（admin）
func (s *UserService) List(page, size int) ([]*models.User, int64, error) {
	if size <= 0 {
		size = 10
	}
	return s.userRepo.List(page, size)
}

// GetUserStats 用户统计
func (s *UserService) GetUserStats() (*models.StatsResult, error) {
	total, _ := s.userRepo.CountTotal()
	todayNew, _ := s.userRepo.CountTodayNew()
	todayActive, _ := s.userRepo.CountTodayActive()
	return &models.StatsResult{
		TotalCount:    total,
		TodayNewCount: todayNew,
		TodayActive:   todayActive,
	}, nil
}
