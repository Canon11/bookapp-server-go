package router

import (
	"bookapp/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.BookList)
	router.POST("/books", controllers.BookCreate)
	router.GET("/books/:id", controllers.BookDetail)
	router.PATCH("/books/:id", controllers.BookEdit)
	router.DELETE("/books/:id", controllers.BookDelete)

	router.GET("/categories", nil)
	router.POST("/categories", nil)
	router.GET("/categories/:id", nil)
	router.PATCH("/categories/:id", nil)
	router.DELETE("/categories/:id", nil)

	return router
}
