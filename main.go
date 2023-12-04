package main

import (
	"be-go-bookshelf/app"
	"be-go-bookshelf/auth"
	"be-go-bookshelf/db"
	"be-go-bookshelf/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.InitDB()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	handler := app.New(db)

	r.GET("/", auth.HomeHandler)
	r.GET("/login", auth.LoginGetHandler)

	r.GET("/books", middleware.AuthValidator, handler.GetBooks)
	r.GET("/book/:id", middleware.AuthValidator, handler.GetBookById)

	r.GET("/addBook", middleware.AuthValidator, handler.GetAddBook)
	r.POST("/book", middleware.AuthValidator, handler.PostBook)

	r.GET("/updateBook/:id", middleware.AuthValidator, handler.GetUpdateBook)
	// the template engine HTML approach doesn't support PUT
	r.POST("/updateBook/:id", middleware.AuthValidator, handler.PutBook)

	// the template engine HTML approach doesn't support DELETE
	r.POST("/deleteBook/:id", middleware.AuthValidator, handler.DeleteBook)

	r.POST("/login", auth.LoginPostHandler)

	r.Run()
}
