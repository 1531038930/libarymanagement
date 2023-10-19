package controller

import (
	"github.com/gin-gonic/gin"
	"librarymanagement/model"
	"net/http"
)

const (
	registhtml = "regist.html"
)

type Regist struct {
}

func (Regist Regist) Static(c *gin.Context) {
	c.HTML(http.StatusOK, registhtml, gin.H{})
}
func (Regist Regist) Registing(c *gin.Context) {
	var regUser model.User
	if err := c.ShouldBind(&regUser); err != nil { //数据绑定失败
		c.JSON(http.StatusOK, gin.H{
			"msg":  "注册失败！",
			"user": regUser,
		})
	} else {
		err = regUser.Add()
		if err != nil { //sql添加失败
			c.JSON(http.StatusOK, gin.H{
				"msg":  "注册失败！",
				"user": regUser,
				"err":  err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "注册成功！",
				"user": regUser,
				"err":  err,
			})
		}
	}
}
