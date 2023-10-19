package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cookie struct {
}

func (Cookie) Get(c *gin.Context) { //没有从cookie获取到id就重新登录
	userid, _ := c.Cookie("userid")
	if len(userid) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "请重新登录",
		})
		c.Abort()
	}
}
