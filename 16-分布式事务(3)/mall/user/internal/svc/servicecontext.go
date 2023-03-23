package svc

import (
	"database/sql"
	"user/database"
	"user/internal/config"
	"user/internal/dao"
	"user/internal/repo"
)

type ServiceContext struct {
	Config   config.Config
	UserRepo repo.UserRepo
	DB       *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	connect := database.Connect(c.MySQL.DataSource, c.CacheRedis)
	db, _ := connect.Conn.RawDB()
	return &ServiceContext{
		Config:   c,
		UserRepo: dao.NewUserDao(connect),
		DB:       db,
	}
}
