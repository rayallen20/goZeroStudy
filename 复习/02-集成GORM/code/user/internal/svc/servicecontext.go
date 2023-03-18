package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
	"user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	port := strconv.Itoa(c.MySQL.Port)
	dsn := c.MySQL.User + ":" + c.MySQL.Password + "@(" + c.MySQL.Host + ":" + port + ")/" + c.MySQL.Database + "?charset=" + c.MySQL.Charset + "&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 禁用复数表名
			SingularTable: true,
		},
	})
	if err != nil {
		panic("连接MySQL失败: " + err.Error())
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
