package model

import "database/sql"

type User struct {
	Id     int64          `db:"id"`
	Name   sql.NullString `db:"name"`
	Gender sql.NullString `db:"gender"`
}

func (u *User) TableName() string {
	return "user"
}
