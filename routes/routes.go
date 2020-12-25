package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User ユーザーモデル
type User struct {
	ID   int    `json:"id" sql:"AUTO_INCREMENT"`
	Name string `json:"name"`
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "docker"
	PASS := "docker"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "sample_gin_dev"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// Home root画面
func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{"Title": "index!!"})
}

// Login ログイン
func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

// SignUp サインアップ
func SignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", gin.H{})
}

// MessagePosts 投稿画面
func MessagePosts(ctx *gin.Context) {
	db := gormConnect()
	defer db.Close()
	user := User{}
	var users []User
	db.Find(&users)
	ctx.HTML(http.StatusOK, "message_posts.html", gin.H{"Users": users, "name": user.Name})
}

// MessageCreate 保存画面
func MessageCreate(ctx *gin.Context) {
	db := gormConnect()
	defer db.Close()
	user := User{}
	user.Name = ctx.PostForm("name")
	db.Create(&user)
	ctx.Redirect(http.StatusSeeOther, "/message_posts")
}

// NoRoute not found画面
func NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
