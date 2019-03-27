package router

import (
	"bookapp/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))

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
