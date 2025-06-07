package main

import (
	"os"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/image-resizer/1/config"
	"github.com/image-resizer/1/controllers"
	_ "github.com/image-resizer/1/docs"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/joho/godotenv"
)

// @title Image Resizer API
// @version 1.0
// @description Resize images easily
func main() {
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, relying on environment variables")
    }
	config.InitDB()
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/resize", controllers.ResizeImageAPI)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}