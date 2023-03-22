package svc

import (
	"user/database"
	"user/internal/config"
	"user/internal/dao"
	"user/internal/repo"
)

type ServiceContext struct {
	Config   config.Config
	UserRepo repo.UserRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserRepo: dao.NewUserDao(database.Connect(c.MySQL.DataSource, c.CacheRedis)),
	}
}
