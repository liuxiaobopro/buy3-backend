package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"user_srv/global"
	"user_srv/model"
)

func DB() {
	// dsn := "abc:123456@tcp(127.0.0.1:3306)/mxshop_user_srv?charset=utf8&parseTime=True&loc=Local"
	m := global.Conf.MysqlInfo
	// fmt.Println(m)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.Name)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	if err := global.DB.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}
}
