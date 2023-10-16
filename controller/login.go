package controller

import (
	"github.com/gin-gonic/gin"
	"librarymanagement/model"
	"net/http"
	"strconv"
)

const (
	loginhtml = "login.html"
)

type Login struct {
}

func (Login Login) Static(c *gin.Context) {
	c.HTML(http.StatusOK, loginhtml, gin.H{})
}

func (Login Login) Loging(c *gin.Context) {
	uname := c.PostForm("username")
	pwd := c.PostForm("password")
	user := &model.User{UserName: uname, Pwd: pwd}
	err := user.Check()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err,
			"msg": "用户名或密码错误！",
		})
	} else {
		c.SetCookie("userid", strconv.FormatUint(uint64(user.UserId), 10), 6000, "/", "localhost", false, true)
		c.String(http.StatusOK, "登录成功")
	}
}
