package logic

import (
	"goginProject/models"
	"testing"
)

func TestSignup(t *testing.T) {
	p := &models.ParamSignUp{
		UserName:   "wbn",
		Password:   "123456",
		RePassword: "123456",
	}

	Signup(p)
}
