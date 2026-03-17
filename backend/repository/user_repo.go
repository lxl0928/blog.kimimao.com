package repository

import (
	"database/sql"

	"blog.kimimao.com/backend/models"
)

// UserRepository 用户数据访问
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository 创建用户仓库
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create 创建用户
func (r *UserRepository) Create(u *models.User) error {
	return r.db.QueryRow(
		`INSERT INTO tb_user (username, real_name, password, phone, email, intro, avatar_url, role)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at, updated_at`,
		u.Username, u.RealName, u.Password, u.Phone, u.Email, u.Intro, u.AvatarURL, u.Role,
	).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
}

// GetByID 根据ID获取用户
func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	u := &models.User{}
	err := r.db.QueryRow(
		`SELECT id, username, real_name, password, phone, email, intro, avatar_url, role, created_at, updated_at, last_login_at
		 FROM tb_user WHERE id = $1`, id,
	).Scan(&u.ID, &u.Username, &u.RealName, &u.Password, &u.Phone, &u.Email, &u.Intro, &u.AvatarURL, &u.Role,
		&u.CreatedAt, &u.UpdatedAt, &u.LastLoginAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return u, err
}

// GetByUsername 根据用户名获取
func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	u := &models.User{}
	err := r.db.QueryRow(
		`SELECT id, username, real_name, password, phone, email, intro, avatar_url, role, created_at, updated_at, last_login_at
		 FROM tb_user WHERE username = $1`, username,
	).Scan(&u.ID, &u.Username, &u.RealName, &u.Password, &u.Phone, &u.Email, &u.Intro, &u.AvatarURL, &u.Role,
		&u.CreatedAt, &u.UpdatedAt, &u.LastLoginAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return u, err
}

// Update 更新用户
func (r *UserRepository) Update(u *models.User) error {
	_, err := r.db.Exec(
		`UPDATE tb_user SET real_name=$1, phone=$2, email=$3, intro=$4, avatar_url=$5, updated_at=CURRENT_TIMESTAMP
		 WHERE id=$6`, u.RealName, u.Phone, u.Email, u.Intro, u.AvatarURL, u.ID,
	)
	return err
}

// UpdatePassword 更新密码
func (r *UserRepository) UpdatePassword(id int64, password string) error {
	_, err := r.db.Exec(`UPDATE tb_user SET password=$1, updated_at=CURRENT_TIMESTAMP WHERE id=$2`, password, id)
	return err
}

// UpdateLastLogin 更新最后登录时间
func (r *UserRepository) UpdateLastLogin(id int64) error {
	_, err := r.db.Exec(`UPDATE tb_user SET last_login_at=CURRENT_TIMESTAMP WHERE id=$1`, id)
	return err
}

// List 分页列表
func (r *UserRepository) List(page, size int) ([]*models.User, int64, error) {
	var total int64
	r.db.QueryRow(`SELECT COUNT(*) FROM tb_user`).Scan(&total)

	offset := (page - 1) * size
	rows, err := r.db.Query(
		`SELECT id, username, real_name, phone, email, intro, avatar_url, role, created_at, last_login_at
		 FROM tb_user ORDER BY id DESC LIMIT $1 OFFSET $2`, size, offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []*models.User
	for rows.Next() {
		u := &models.User{}
		var lastLogin sql.NullTime
		err := rows.Scan(&u.ID, &u.Username, &u.RealName, &u.Phone, &u.Email, &u.Intro, &u.AvatarURL, &u.Role,
			&u.CreatedAt, &lastLogin)
		if err != nil {
			return nil, 0, err
		}
		if lastLogin.Valid {
			u.LastLoginAt = &lastLogin.Time
		}
		list = append(list, u)
	}
	return list, total, nil
}

// CountTotal 用户总数
func (r *UserRepository) CountTotal() (int64, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM tb_user`).Scan(&count)
	return count, err
}

// CountTodayNew 今日新增
func (r *UserRepository) CountTodayNew() (int64, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM tb_user WHERE DATE(created_at) = CURRENT_DATE`).Scan(&count)
	return count, err
}

// CountTodayActive 今日活跃
func (r *UserRepository) CountTodayActive() (int64, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM tb_user WHERE DATE(last_login_at) = CURRENT_DATE`).Scan(&count)
	return count, err
}
