package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/ltinyho/lt-go-project/app/user/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSetData = wire.NewSet(NewDB, NewData, NewUserRepo)

type Data struct {
	db *gorm.DB
}

func NewDB(conf *conf.Data) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&user{})
	if err != nil {
		panic(err)
	}
	return db, nil
}

func NewData(db *gorm.DB) (*Data, func(), error) {
	d := &Data{db: db}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	return d, func() {
		sqlDb.Close()
	}, nil
}
