package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"goginProject/models"
)

//数据访问层
const secret = "rochar"

var (
	ErrorUserExist = errors.New("用户已经存在")
)

//用户校验
func CheckUserExist(username string) error {
	sql := `select count(user_id) from user where username=?`
	var count int64
	if err := db.Get(&count, sql, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

//加密
func encryptPassword(opassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	
	return hex.EncodeToString(h.Sum([]byte(opassword)))
}

//插入用户
func InsertUser(user *models.User) error {
	user.Password = encryptPassword(user.Password)

	sql := `insert into user (user_id,username,password) values (?,?,?)`

	_, err := db.Exec(sql, user.UserId, user.UserName, user.Password)
	if err != nil {
		zap.L().Error("InsertUser error", zap.Error(err))
		fmt.Println(err)
	}
	return err
}
