package controllers

import (
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary hello worldを返す
// @Schemes
// @Description Hello Worldを返すだけのテスト用
// @Tags Hello World
// @Accept json
// @Produce json
// @Success 200 {string} Hello world
// @Router /hello [get]
func GetHelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
