package server

import (
	"github.com/gin-gonic/gin"
)

func ListenAndServe() {
	// Create a new Gin router
	router := gin.Default()

	// Define endpoints
	router.GET("/flows", GetAllFlows)
	router.GET("/flows/:id", GetFlowByID)
	router.POST("/flows", CreateFlow)

	// Start the server on port 8080
	router.Run(":8080")
}
