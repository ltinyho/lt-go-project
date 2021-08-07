package biz

import "context"

type User struct {
	Uid  int64
	Name string
	Username string
	Password string
	Age  int64
}

type UserRepo interface {
	Create(ctx context.Context, u *User) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	out, err := uc.repo.Create(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}
