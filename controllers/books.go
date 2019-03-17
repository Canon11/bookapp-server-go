package controllers

import (
	"github.com/gin-gonic/gin"
)

type testResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func BookList(c *gin.Context) {
	res := testResp{ID: 2, Name: "book_list"}
	c.JSON(200, res)
}

func BookCreate(c *gin.Context) {
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
