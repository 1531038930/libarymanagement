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
	if err := c.ShouldBind(&newBook); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "添加失败！",
			"user": newBook,
		})
	} else {
		err = newBook.Add()
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
func (Borrow) Info(c *gin.Context) {
	bookid, err := strconv.Atoi(c.Param("bookid"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "bookid错误",
			"err": err,
		})
		c.Abort()
	}
	var list model.B_list
	list.Bookid = uint(bookid)
	useridstr, err := c.Cookie("userid")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "userid错误",
			"err": err,
		})
		c.Abort()
	}
	userid, err := strconv.Atoi(useridstr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "userid错误",
			"err": err,
		})
		c.Abort()
	}
	list.UserId = uint(userid)
	err = list.Add()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "借书错误",
			"err": err,
		})
		c.Abort()
	}
}

/*
func (one Borrow) Info(c *gin.Context) {
	bookid, err := strconv.Atoi(c.Param("bookid"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "bookid错误",
			"err": err,
		})
		c.Abort()
	}
	one.Book.Bookid = uint(bookid)
	useridstr, err := c.Cookie("userid")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "userid错误",
			"err": err,
		})
		c.Abort()
	}
	userid, err := strconv.Atoi(useridstr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "userid错误",
			"err": err,
		})
		c.Abort()
	}
	one.User.UserId = uint(userid)

}
*/
