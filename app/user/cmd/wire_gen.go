// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/ltinyho/lt-go-project/app/user/internal/biz"
	"github.com/ltinyho/lt-go-project/app/user/internal/conf"
	"github.com/ltinyho/lt-go-project/app/user/internal/data"
	"github.com/ltinyho/lt-go-project/app/user/internal/server"
	"github.com/ltinyho/lt-go-project/app/user/internal/service"
	"github.com/ltinyho/lt-go-project/pkg/gapp"
)

// Injectors from wire.go:

func initApp(confServer *conf.Server, confData *conf.Data) (*gapp.App, func(), error) {
	db, err := data.NewDB(confData)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(db)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData)
	userUseCase := biz.NewUserUseCase(userRepo)
	userService := service.NewUserService(userUseCase)
	httpServer := server.NewHttpServer(confServer, userService)
	app := newApp(httpServer)
	return app, func() {
		cleanup()
	}, nil
}
