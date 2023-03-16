package database

import "github.com/zeromicro/go-zero/core/stores/sqlx"

type DBConn struct {
	Conn sqlx.SqlConn
}

func Connect(dataSource string) *DBConn {
	return &DBConn{Conn: sqlx.NewMysql(dataSource)}
}
