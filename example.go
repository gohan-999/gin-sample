package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gohan-999/gin-sample/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	router.GET("/", routes.Home)
	router.GET("/message_posts", routes.MessagePosts)
	router.POST("/message_create", routes.MessageCreate)
	router.GET("/login", routes.Login)
	router.GET("/signup", routes.SignUp)
	router.NoRoute(routes.NoRoute)

	router.Run()
}
