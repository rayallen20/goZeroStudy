package database

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DBConn struct {
	Conn      sqlx.SqlConn
	ConnCache sqlc.CachedConn
}

func Connect(dataSource string, conf cache.CacheConf) *DBConn {
	sqlConn := sqlx.NewMysql(dataSource)
	d := &DBConn{
		Conn: sqlConn,
	}

	// redis配置不为空 说明需要用缓存 再连接缓存
	if conf != nil {
		cachedConn := sqlc.NewConn(sqlConn, conf)
		d.ConnCache = cachedConn
	}
	return d
}
