package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"librarymanagement/model"
	"librarymanagement/toredis"
	"net/http"
	"strconv"
)

type Cookie struct {
	token  string
	userid uint
}

func (me *Cookie) Getcookie(c *gin.Context) error {
	idcookie, err := c.Cookie("userid")
	if err != nil {
		return errors.New("未获取到id")
	}
	id, err := strconv.Atoi(idcookie)
	if err != nil {
		return errors.New("id错误")
	}
	me.userid = uint(id)
	tokencookie, err := c.Cookie("token")
	if err != nil {
		return errors.New("未获取到token")
	}
	me.token = tokencookie
	return nil
}
func (Cookie) Get(c *gin.Context) { //没有获取到token、id或不匹配就重新登录
	onecookie := &Cookie{}
	if err := onecookie.Getcookie(c); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "cookie获取失败，请重新登录",
		})
		c.Abort()
		return
	}
	err := toredis.Check(onecookie.token, fmt.Sprint(onecookie.userid))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		c.Abort()
	}
}
func (Cookie) GetWithLevel(c *gin.Context) {
	onecookie := &Cookie{}
	if err := onecookie.Getcookie(c); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "cookie获取失败，请重新登录",
		})
		c.Abort()
		return
	}
	if err := toredis.Check(onecookie.token, fmt.Sprint(onecookie.userid)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		c.Abort()
		return
	}
	var ckUser model.User
	ckUser.UserId = onecookie.userid //需要level=1
	if err := ckUser.CheckLevel(); err != nil {
		c.String(http.StatusOK, "权限不足！")
		c.Abort()
	}
}
