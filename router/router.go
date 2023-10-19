package router

import (
	"github.com/gin-gonic/gin"
	"librarymanagement/controller"
)

const (
	htmlpath = "./static/html/*"
	csspath  = "./static/css"
)

func LMsys() {
	r := gin.Default()
	r.LoadHTMLGlob(htmlpath)
	r.Static("/css", csspath)
	regist := r.Group("/regist") //注册
	{
		regist.GET("/", controller.Regist{}.Static)
		regist.POST("/", controller.Regist{}.Registing)
	}
	login := r.Group("/login") //登录
	{
		login.GET("/", controller.Login{}.Static)
		login.POST("/", controller.Login{}.Loging)
	}
	book := r.Group("/book") //增加书籍、借书、归还
	{
		book.GET("/", controller.Cookie{}.Get, controller.Borrow{}.Static)
		book.POST("/add", controller.Borrow{}.Add)                                       //添加书
		book.GET("/borrow/:bookid", controller.Cookie{}.Get, controller.Borrow{}.Info)   //借书
		book.GET("/return/:bookid", controller.Cookie{}.Get, controller.Borrow{}.Revert) //归还
	}
	r.Run()
}
