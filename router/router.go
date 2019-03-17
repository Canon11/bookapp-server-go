package router

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/books", nil)
	router.POST("/books", nil)
	router.GET("/books/:id", nil)
	router.PATCH("/books/:id", nil)
	router.DELETE("/books/:id", nil)

	router.GET("/categories", nil)
	router.POST("/categories", nil)
	router.GET("/categories/:id", nil)
	router.PATCH("/categories/:id", nil)
	router.DELETE("/categories/:id", nil)

	return router
}
