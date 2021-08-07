package service

import (
	"github.com/google/wire"
	v1 "github.com/ltinyho/lt-go-project/api/user/v1"
	"github.com/ltinyho/lt-go-project/app/user/internal/biz"
)

var ProviderSetService = wire.NewSet(NewUserService)

type UserService struct {
	v1.UnimplementedUserServer
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}
