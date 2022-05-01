package mysql

import (
	"fmt"

	config "Apale7/simple_object_storage/config_loader"
	mysql_model "Apale7/simple_object_storage/dal/mysql/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Option func(*gorm.DB) *gorm.DB

const (
	dsnParttern = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

var db *gorm.DB

func Init() {
	var err error
	dsn := fmt.Sprintf(dsnParttern, config.Get("mysql_user"), config.Get("mysql_password"), config.Get("mysql_host"), config.GetInt("mysql_port"), config.Get("mysql_db"))
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		CreateBatchSize:                          200,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	CreateTable()
}

func getDB() *gorm.DB {
	return db.Session(&gorm.Session{NewDB: true}).Debug()
}

func CreateTable() {
	err := db.AutoMigrate(&mysql_model.FileMeta{}, &mysql_model.FileLink{})
	if err != nil {
		panic(err)
	}
}
