package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tetn39/MelodyStatus/controllers"
)

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

	r.Run()

}
