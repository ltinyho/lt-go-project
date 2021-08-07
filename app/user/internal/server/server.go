package server

import "github.com/google/wire"

// ProviderSetServer ProviderSet 初始化 server
var ProviderSetServer = wire.NewSet(NewHttpServer)
