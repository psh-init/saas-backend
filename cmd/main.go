package main

import (
	"log"
	"os"
	"saas-backend/internal"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	internal.ConnectDB()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "SaaS Backend is running!",
		})
	})

	port := os.Getenv("PORT")
	log.Println("Server starting on port:", port)
	r.Run(":" + port)
}
