package user

import (
	"context"
	"t-blog-back/models"
	"t-blog-back/pkg/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// LoginForm 登录form
type LoginForm struct {
	UserName string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
}

// LoginData 登录数据
type LoginData struct {
	Token string `json:"token"`
}

// Login 登录
func Login(ctx context.Context, loginForm LoginForm) (lData LoginData, err error) {
	user := models.User{}
	if err = user.FindUserByName(ctx, loginForm.UserName); err == nil && bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(loginForm.PassWord)) == nil {
		token := utils.GenToken()
		if _, err = user.UpdateToken(ctx, token); err == nil {
			lData.Token = token
		}
	}
	return
}

// InitUserForm form
type InitUserForm struct {
}

// InitUserData data
type InitUserData struct {
	UserName string `json:"userName"`
	UserPass string `json:"userPass"`
}

// InitUser 初始化一个用户
func InitUser(ctx context.Context, iForm InitUserForm) (iData InitUserData, err error) {
	now := time.Now()
	user := models.User{}
	initName := "admin"
	initPass := "admin"

	iData.UserName = initName
	iData.UserPass = initPass

	if err = user.FindUserByName(ctx, initName); err != nil {
		pass := []byte{}
		if pass, err = bcrypt.GenerateFromPassword([]byte(initPass), 16); err == nil {
			user.AuthorID = utils.GenUniqueID()
			user.Pass = string(pass)
			if user.Name == "" {
				user.Name = initName
				user.AddTime = now
				user.ModifyTime = now
				_, err = user.AddUser(ctx)
			} else {
				user.ModifyTime = now
				_, err = user.UpdateUser(ctx)
			}
		}
	}
	return
}
