package svc

import (
	"user-score/database"
	"user-score/internal/config"
	"user-score/internal/dao"
	"user-score/internal/repo"
)

type ServiceContext struct {
	Config        config.Config
	UserScoreRepo repo.UserScoreRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserScoreRepo: dao.NewUserScoreDao(database.Connect(c.MySQL.DataSource, c.CacheRedis)),
	}
}
