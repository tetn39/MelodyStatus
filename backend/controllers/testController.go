package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tetn39/MelodyStatus/models"
)

// GetHelloWorldはHello Worldを返すエンドポイント
// PingExample godoc
// @Summary hello worldを返す
// @Schemes
// @Description Hello Worldを返すだけのテスト用
// @Tags Hello World
// @Accept json
// @Produce json
// @Success 200 {object} models.HelloStruct
// @Router /hello [get]
func GetHelloWorld(c *gin.Context) {
	var hello models.HelloStruct
	hello.Message = "Hello World"

	c.JSON(200, hello)
}
