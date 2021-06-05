package logic

import (
	"goginProject/dao/mysql"
	"goginProject/models"
	"goginProject/pkg/snowflake"
)

//业务层

func Signup(p *models.ParamSignUp) error {
	//判断用户是否存在
	err := mysql.CheckUserExist(p.UserName)
	if err != nil {
		return err
	}
	//生成雪花ID
	userID := snowflake.GenID()
	user := &models.User{
		UserId:   userID,
		UserName: p.UserName,
		Password: p.Password,
	}
	err = mysql.InsertUser(user)
	return err
}
