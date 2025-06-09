package main

import "github.com/gin-gonic/gin"

func main() {
	// rb := models.Rabbitmq{}
	// rb.InitRabbitMQ()
	// rb.CloseRabbitMQ()

	// database.Connect()
	// database.Migrate()

	// startBackgroundWorkers()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run("0.0.0.0:8080")
}
