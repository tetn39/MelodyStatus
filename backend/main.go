package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tetn39/MelodyStatus/controllers"
	"github.com/tetn39/MelodyStatus/docs"
	"github.com/tetn39/MelodyStatus/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDb()
	initializers.SpotifyOAuth()
}

// @title MelodyStatusのAPI一覧
// @version 1.0
// @license.name tetn39
// @description MelodyStatusのAPI一覧だよ
func main() {
	route := SetupRoutes()
	docs.SwaggerInfo.BasePath = "/api/v1"
	fmt.Println("起動しました")

	v1 := route.Group("/api/v1")
	{
		// エンドポイントの設定
		v1.GET("/hello", controllers.GetHelloWorld)                      // Hello Worldエンドポイント
		v1.GET("/musics", controllers.GetAllMusics)                      // 音楽データの取得エンドポイント
		v1.POST("/musics", controllers.PostMusic)                        // 音楽データの登録エンドポイント
		v1.GET("/auth/spotify/login", controllers.GetSpotifyLogin)       // Spotifyログインエンドポイント
		v1.GET("/auth/spotify/callback", controllers.GetSpotifyCallback) // Spotifyコールバックエンドポイント
	}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	route.Run(":8080")
}

// SetupRoutesはGinのルーティングをセットアップする関数
func SetupRoutes() *gin.Engine {
	route := gin.Default()

	// CORS設定
	route.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	// ルートエンドポイントの設定
	route.GET("/", controllers.GetHelloWorld)

	return route
}
