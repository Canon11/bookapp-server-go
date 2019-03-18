package controllers

import (
	"bookapp/models"
	"log"

	"github.com/gin-gonic/gin"
)

type testResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func BookList(c *gin.Context) {
	client, err := models.NewClient()
	if err != nil {
		c.Error(err)
	}

	books, err := client.ListBook(c)
	log.Printf("%+v", books)
	if err != nil {
		c.Error(err)
	}

	c.JSON(200, books)
}

func BookCreate(c *gin.Context) {
	client, err := models.NewClient()
	if err != nil {
		c.Error(err)
	}

	book := models.Book{
		Name:     "TestBook",
		Category: 1,
	}
	if err := client.CreateBook(c, &book); err != nil {
		c.Error(err)
	}

	res := testResp{ID: 3, Name: "book_create"}
	c.JSON(200, res)
}

func BookDetail(c *gin.Context) {
	res := testResp{ID: 4, Name: "book_detail"}
	c.JSON(200, res)
}

func BookEdit(c *gin.Context) {
	res := testResp{ID: 5, Name: "book_edit"}
	c.JSON(200, res)
}

func BookDelete(c *gin.Context) {
	res := testResp{ID: 6, Name: "book_delete"}
	c.JSON(200, res)
}
