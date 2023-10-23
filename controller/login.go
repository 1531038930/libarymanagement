package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"librarymanagement/model"
	"librarymanagement/toredis"
	"net/http"
	"strconv"
	"time"
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
	user := &model.User{UserName: uname, Pwd: pwd} //账号密码赋值一个实例
	if err := user.Check(); err != nil {           //实例进行账号密码校验
		c.JSON(http.StatusOK, gin.H{
			"err": err,
			"msg": "用户名或密码错误！",
		})
		return
	} else { //校验通过则设置token
		src := []byte(strconv.Itoa(int(time.Now().UnixNano()))) //获得当前时间戳
		//fmt.Println(src, "--")
		tok := jwt.New(jwt.SigningMethodHS256) //加密算法
		token, err := tok.SignedString(src)    //生成token
		if err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}
		//fmt.Println(token, err)
		if err := toredis.Add(user.UserId, token); err != nil {
			fmt.Println("redis连接失败")
			c.String(http.StatusOK, err.Error())
			return
		}
		c.SetCookie("token", token, 6000, "/", "localhost", false, true)
		c.SetCookie("userid", strconv.FormatUint(uint64(user.UserId), 10), 6000, "/", "localhost", false, true)
		c.String(http.StatusOK, "登录成功", user.UserId)
	}
}
