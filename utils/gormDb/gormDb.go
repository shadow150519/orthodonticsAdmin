package gormDb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hello/orthodonticsAdmin/global/my_errors"
	"hello/orthodonticsAdmin/global/variable"
	"log"
	"time"
)

func InitGormDbMysql(){
	port := variable.ConfigViper.GetString("mysql.port")
	username := variable.ConfigViper.GetString("mysql.username")
	password := variable.ConfigViper.GetString("mysql.password")
	database := variable.ConfigViper.GetString("mysql.database")
	host := variable.ConfigViper.GetString("mysql.host")

	dsn := username +":"+password+ "@tcp(" + host + ":" + port + ")/" + database +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	variable.GormDbMysql, err = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal(my_errors.ErrorsGormInitFail + err.Error())
	}
	db, _ := variable.GormDbMysql.DB()
	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(128)
}
