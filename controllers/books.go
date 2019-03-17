package controllers

import (
	"github.com/gin-gonic/gin"
)

type testResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func BookIndex(c *gin.Context) {
	res := testResp{ID: 2, Name: "book_index"}
	c.JSON(200, res)
}
