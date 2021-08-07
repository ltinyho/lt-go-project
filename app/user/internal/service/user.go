package service

import (
	"context"
	v1 "github.com/ltinyho/lt-go-project/api/user/v1"
	"github.com/ltinyho/lt-go-project/app/user/internal/biz"
)

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserRes,
	error) {
	rv, err := s.uc.Create(ctx, &biz.User{
		Name: req.Name,
		Age: int64(req.Age),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserRes{
		Uid:      rv.Uid,
		Username: rv.Username,
	}, nil
}
