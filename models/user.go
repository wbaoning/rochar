package models

//对应数据库实体的ID

type User struct {
	UserId   int64  `db:"user_id"`
	UserName string `db:"username"`
	Password string `db:"password"`
}
