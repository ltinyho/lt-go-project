// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ltinyho/lt-go-project/app/user/internal/biz"
	"github.com/ltinyho/lt-go-project/app/user/internal/conf"
	"github.com/ltinyho/lt-go-project/app/user/internal/data"
	"github.com/ltinyho/lt-go-project/app/user/internal/server"
	"github.com/ltinyho/lt-go-project/app/user/internal/service"
	"github.com/ltinyho/lt-go-project/pkg/gapp"
)

func initApp(confServer *conf.Server, confData *conf.Data) (*gapp.App, func(), error) {
	panic(wire.Build(
		data.ProviderSetData,
		biz.ProviderSetBiz,
		service.ProviderSetService,
		server.ProviderSetServer,
		newApp,
	))
}
