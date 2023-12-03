package main

import (
	"be-go-bookshelf/app"
	"be-go-bookshelf/auth"
	"be-go-bookshelf/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.InitDB()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	handler := app.New(db)

	r.GET("/", auth.HomeHandler)
	r.GET("/login", auth.LoginGetHandler)

	r.GET("/books", handler.GetBooks)

	r.POST("/login", auth.LoginPostHandler)

	r.Run()
}
