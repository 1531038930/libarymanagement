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
	regist := r.Group("/regist")
	{
		regist.GET("/", controller.Regist{}.Static)
		regist.POST("/", controller.Regist{}.Registing)
	}
	login := r.Group("/login")
	{
		login.GET("/", controller.Login{}.Static)
		login.POST("/", controller.Login{}.Loging)
	}
	book := r.Group("/book")
	{
		book.GET("/", controller.Cookie{}.Get, controller.Borrow{}.Static)
		book.POST("/add", controller.Borrow{}.Add)                                     //添加书
		book.GET("/borrow/:bookid", controller.Cookie{}.Get, controller.Borrow{}.Info) //借书
		book.GET("/return/:bookid")
	}
	r.Run()
}
