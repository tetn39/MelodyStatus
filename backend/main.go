package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tetn39/MelodyStatus/controllers"
	"github.com/tetn39/MelodyStatus/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDb()
}

func main() {
	fmt.Println("Hello World")
	r := gin.Default()
	// CORS設定
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	// hello world
	r.GET("/api/v1/hello", controllers.GetHelloWorld)

	// 音楽データの取得
	r.GET("/api/v1/musics", controllers.GetAllMusics)

	// 音楽データの登録
	r.POST("/api/v1/musics", controllers.PostMusic)
	r.Run()

}
