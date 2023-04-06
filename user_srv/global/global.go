package global

import (
	"gorm.io/gorm"

	"user_srv/config"
)

var (
	DB   *gorm.DB
	Conf config.Conf
)
