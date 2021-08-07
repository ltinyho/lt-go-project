package server

import (
	v1 "github.com/ltinyho/lt-go-project/api/user/v1"
	"github.com/ltinyho/lt-go-project/app/user/internal/conf"
	"github.com/ltinyho/lt-go-project/app/user/internal/service"
	"github.com/ltinyho/lt-go-project/pkg/gapp"
)

func NewHttpServer(c *conf.Server,s *service.UserService) *gapp.HttpServer {
	srv:= gapp.NewServer(gapp.Addr(c.Http.Addr))
	v1.RegisterUserServer(srv,s)
	return srv
}
