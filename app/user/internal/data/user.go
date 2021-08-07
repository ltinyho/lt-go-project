package data

import (
	"context"
	"github.com/ltinyho/lt-go-project/app/user/internal/biz"
	"gorm.io/gorm"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data, ) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

type user struct {
	gorm.Model
	Name     string
	Username string
	Password string
	Age      int64
}

func (r *userRepo) Create(ctx context.Context, b *biz.User) (*biz.User, error) {
	u := user{
		Name:     b.Name,
		Username: b.Username,
		Password: b.Password, // 加密
		Age:      b.Age,
	}
	err := r.data.db.WithContext(ctx).Create(&u).Error
	return &biz.User{
		Uid:      int64(u.ID),
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		Age:      u.Age,
	}, err
}
