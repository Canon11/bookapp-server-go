package controllers

import (
	"github.com/gin-gonic/gin"
)

func CategoryList(c *gin.Context) {
	res := testResp{ID: 2, Name: "category_list"}
	c.JSON(200, res)
}

func CategoryCreate(c *gin.Context) {
	res := testResp{ID: 3, Name: "category_create"}
	c.JSON(200, res)
}

func CategoryDetail(c *gin.Context) {
	res := testResp{ID: 4, Name: "category_detail"}
	c.JSON(200, res)
}

func CategoryEdit(c *gin.Context) {
	res := testResp{ID: 5, Name: "category_edit"}
	c.JSON(200, res)
}

func CategoryDelete(c *gin.Context) {
	res := testResp{ID: 6, Name: "category_delete"}
	c.JSON(200, res)
}
