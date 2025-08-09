package main

import (
	"flowgraph/db"
	"flowgraph/flows"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB database client (Panics if connection fails)
	database := db.InitMongoDbClient()

	// DDD hamburger architecture: https://medium.com/@remast/the-ddd-hamburger-for-go-61dba99c4aaf
	flowRepository := flows.NewFlowsRepository(database)
	flowService := flows.NewFlowService(flowRepository)
	flowController := flows.NewFlowRestController(flowService)

	// Start HTTP server
	ListenAndServe(flowController)
}

func ListenAndServe(flowController *flows.FlowRestController) {
	// Create a new Gin router
	router := gin.Default()

	// Mount HTTP server endpoints
	router.GET("/flows", flowController.GetAllFlows)
	router.GET("/flows/:id", flowController.GetFlowByID)
	router.POST("/flows", flowController.CreateFlow)

	// Start the server on port 8080
	router.Run(":8080")
}
