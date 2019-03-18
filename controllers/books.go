package controllers

import (
	"bookapp/models"
	"log"
	"strconv"

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

	category, err := strconv.Atoi(c.PostForm("category"))
	if err != nil {
		c.Error(err)
	}

	book := models.Book{
		Name:     c.PostForm("name"),
		Category: category,
	}
	if err := client.CreateBook(c, &book); err != nil {
		c.Error(err)
	}

	c.JSON(200, nil)
}

func BookDetail(c *gin.Context) {
	client, err := models.NewClient()
	if err != nil {
		c.Error(err)
	}

	paramID := c.Param("id")
	bookID, err := strconv.Atoi(paramID)
	if err != nil {
		c.Error(err)
	}

	book, err := client.GetBook(c, int64(bookID))
	if err != nil {
		c.Error(err)
	}

	c.JSON(200, book)
}

func BookEdit(c *gin.Context) {
	client, err := models.NewClient()
	if err != nil {
		c.Error(err)
	}

	category, err := strconv.Atoi(c.PostForm("category"))
	if err != nil {
		c.Error(err)
	}

	paramID := c.Param("id")
	bookID, err := strconv.Atoi(paramID)
	if err != nil {
		c.Error(err)
	}

	book := models.Book{
		ID:       int64(bookID),
		Name:     c.PostForm("name"),
		Category: category,
	}
	res, err := client.EditBook(c, &book)
	if err != nil {
		c.Error(err)
	}

	c.JSON(200, res)
}

func BookDelete(c *gin.Context) {
	client, err := models.NewClient()
	if err != nil {
		c.Error(err)
	}

	paramID := c.Param("id")
	bookID, err := strconv.Atoi(paramID)
	if err != nil {
		c.Error(err)
	}

	if err = client.DeleteBook(c, int64(bookID)); err != nil {
		c.Error(err)
	}

	c.JSON(200, nil)
}
