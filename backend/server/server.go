package server

import (
	"flowgraph/db"

	"github.com/gin-gonic/gin"
)

func ListenAndServe() {
	// Create a new Gin router
	router := gin.Default()

	// Define a simple GET endpoint
	router.GET("/hello", func(c *gin.Context) {

		// Test db connection
		db.CreateDocument("test", map[string]interface{}{
			"name":  "test",
			"value": "Hello, World!",
		})

		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
