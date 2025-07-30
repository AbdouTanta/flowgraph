package main

import (
	"github.com/gin-gonic/gin"
)

func ListenAndServe(app *application) {
	// Create a new Gin router
	router := gin.Default()

	// Register endpoints
	router.GET("/flows", app.GetAllFlows)
	router.GET("/flows/:id", app.GetFlowByID)
	router.POST("/flows", app.CreateFlow)

	// Start the server on port 8080
	router.Run(":8080")
}
