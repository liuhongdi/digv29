package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv29/global"
)

type UserController struct{}
func NewUserController() UserController {
	return UserController{}
}
//登录
func (g *UserController) Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	result := global.NewResult(c)

	if (username == "lhd" && password=="123") {
		//set session
		session := sessions.Default(c)
		session.Set("username", username)
		session.Save()
		result.Success("success");
	} else {
		result.Error(1,"用户名密码错误");
	}
	return
}

//读session
func (g *UserController) Session(c *gin.Context) {
    //get session
	session := sessions.Default(c)
	username := session.Get("username")
	result := global.NewResult(c)
	result.Success(username);
	return
}

//logout
func (g *UserController) Logout(c *gin.Context) {
    //clear session
	session := sessions.Default(c)
	session.Clear()
    session.Save()
	result := global.NewResult(c)
	result.Success("success");
	return
}