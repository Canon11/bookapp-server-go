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

	router.GET("/categories", controllers.CategoryList)
	router.POST("/categories", controllers.CategoryCreate)
	router.GET("/categories/:id", controllers.CategoryDetail)
	router.PATCH("/categories/:id", controllers.CategoryEdit)
	router.DELETE("/categories/:id", controllers.CategoryDelete)

	return router
}
