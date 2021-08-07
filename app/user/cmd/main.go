package main

import (
	"fmt"
	"github.com/ltinyho/lt-go-project/app/user/internal/conf"
	"github.com/ltinyho/lt-go-project/pkg/gapp"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	err := parseConf()
	if err != nil {
		panic(err)
	}
	file := viper.ConfigFileUsed()
	fmt.Println("file", file)
	var (
		data   conf.Data
		server conf.Server
	)
	err = viper.UnmarshalKey("data", &data)
	if err != nil {
		panic(err)
	}
	err = viper.UnmarshalKey("server", &server)
	if err != nil {
		panic(err)
	}
	app, cleanup, err := initApp(&server, &data)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}

func newApp(hs *gapp.HttpServer) *gapp.App {
	return gapp.New(
		gapp.Server(hs),
	)
}

func parseConf() (err error) {
	pflag.StringP("conf", "c", "config.yaml", "config file name")

	pflag.Parse()
	err = viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return
	}
	viper.AutomaticEnv()
	// read from config file
	viper.SetConfigFile(viper.GetString("conf"))
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	return nil
}
