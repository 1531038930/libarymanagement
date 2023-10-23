package controller

import (
	"github.com/gin-gonic/gin"
	"librarymanagement/model"
	"net/http"
	"strconv"
)

const (
	borrowhtml = "borrow.html"
)

type Borrow struct {
}

func (Borrow) Static(c *gin.Context) {
	c.HTML(http.StatusOK, borrowhtml, gin.H{})
}
func (Borrow) Add(c *gin.Context) {
	var newBook model.Book
	if err := c.ShouldBind(&newBook); err != nil { //数据绑定失败
		c.JSON(http.StatusOK, gin.H{
			"msg":  "添加失败！",
			"user": newBook,
		})
	} else {
		err = newBook.Add() //向数据库添加书籍信息
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "添加失败！",
				"user": newBook,
				"err":  err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "注册成功！",
				"user": newBook,
				"err":  err,
			})
		}
	}
}
func (Borrow) Info(c *gin.Context) { //借书
	bookid, err := strconv.Atoi(c.Param("bookid")) //将动态路由str参数转化为int
	if err != nil {                                //动态路由转换数字失败
		c.JSON(http.StatusOK, gin.H{
			"msg": "bookid错误",
			"err": err,
		})
		c.Abort()
	}
	var list model.B_list
	list.Bookid = uint(bookid)
	useridstr, err := c.Cookie("userid") //获取userid string
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "userid错误",
			"err": err,
		})
		c.Abort()
		return
	}
	userid, err := strconv.Atoi(useridstr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "userid错误",
			"err": err,
		})
		c.Abort()
		return
	}
	list.UserId = uint(userid) //userid转换为uint类型
	err = list.Add()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "借书错误",
			"err": err,
		})
		c.Abort()
		return
	}
	c.String(http.StatusOK, "借书成功！")
}
func (Borrow) Revert(c *gin.Context) {
	bookid, err := strconv.Atoi(c.Param("bookid")) //取得动态路由参数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "bookid错误",
			"err": err,
		})
		c.Abort()
		return
	}
	var list model.B_list
	list.Bookid = uint(bookid)
	useridstr, err := c.Cookie("userid") //取得用户id
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "userid错误",
			"err": err,
		})
		c.Abort()
		return
	}
	userid, err := strconv.Atoi(useridstr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "userid错误",
			"err": err,
		})
		c.Abort()
		return
	}
	list.UserId = uint(userid)
	err = list.End() //进入归还sql
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "还书错误",
			"err": err.Error(),
		})
		c.Abort()
		return
	} else {
		c.String(http.StatusOK, "还书成功")
	}
}
