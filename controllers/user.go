package controllers

//控制器层

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"goginProject/dao/mysql"
	"goginProject/logic"
	"goginProject/models"
)

func SignUpHandler(c *gin.Context) {
	//获取参数路径
	p := new(models.ParamSignUp)

	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数，直接返回响应
		zap.L().Info("SignUp with invalid param", zap.Error(err))

		err, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
		return
	}

	//业务处理
	err := logic.Signup(p)
	if err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
