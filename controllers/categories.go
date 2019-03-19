package controllers

import (
	"bookapp/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CategoryList(c *gin.Context) {
	client, err := models.NewClient()
	if err != nil {
		c.Error(err)
	}

	categories, err := client.ListCategory(c)
	if err != nil {
		c.Error(err)
	}

	c.JSON(200, categories)
}

func CategoryCreate(c *gin.Context) {
	client, err := models.NewClient()
	if err != nil {
		c.Error(err)
	}

	category := models.Category{
		Name: c.PostForm("name"),
	}
	if err := client.CreateCategory(c, &category); err != nil {
		c.Error(err)
	}
	c.JSON(200, nil)
}

func CategoryDetail(c *gin.Context) {
	client, err := models.NewClient()
	if err != nil {
		c.Error(err)
	}

	paramID := c.Param("id")
	categoryID, err := strconv.Atoi(paramID)
	if err != nil {
		c.Error(err)
	}

	category, err := client.GetBook(c, int64(categoryID))
	if err != nil {
		c.Error(err)
	}

	c.JSON(200, category)
}

func CategoryEdit(c *gin.Context) {
	client, err := models.NewClient()
	if err != nil {
		c.Error(err)
	}

	paramID := c.Param("id")
	categoryID, err := strconv.Atoi(paramID)
	if err != nil {
		c.Error(err)
	}

	category := models.Category{
		ID:   int64(categoryID),
		Name: c.PostForm("name"),
	}
	res, err := client.EditCategory(c, &category)
	if err != nil {
		c.Error(err)
	}

	c.JSON(200, res)
}

func CategoryDelete(c *gin.Context) {
	client, err := models.NewClient()
	if err != nil {
		c.Error(err)
	}

	paramID := c.Param("id")
	categoryID, err := strconv.Atoi(paramID)
	if err != nil {
		c.Error(err)
	}

	if err = client.DeleteCategory(c, int64(categoryID)); err != nil {
		c.Error(err)
	}

	c.JSON(200, nil)
}
