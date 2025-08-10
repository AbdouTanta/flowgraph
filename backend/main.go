package main

import (
	"flowgraph/auth"
	"flowgraph/db"
	"flowgraph/flows"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB database client (Panics if connection fails)
	database := db.InitMongoDbClient()

	// DDD hamburger architecture: https://medium.com/@remast/the-ddd-hamburger-for-go-61dba99c4aaf
	// Flows repository, service, and controller
	flowRepository := flows.NewFlowsRepository(database)
	flowService := flows.NewFlowService(flowRepository)
	flowController := flows.NewFlowRestController(flowService)

	// Auth repository, service, and controller
	authRepository := auth.NewAuthRepository(database)
	authService := auth.NewAuthService(authRepository)
	authController := auth.NewAuthRestController(authService)

	// Start HTTP server
	ListenAndServe(flowController, authController)
}

func ListenAndServe(flowController *flows.FlowRestController, authController *auth.AuthRestController) {
	// Create a new Gin router
	router := gin.Default()

	// Mount HTTP server endpoints
	// Flows endpoints
	router.GET("/flows", flowController.GetAllFlows)
	router.GET("/flows/:id", flowController.GetFlowByID)
	router.POST("/flows", flowController.CreateFlow)
	// Auth endpoints
	router.POST("/login", authController.Login)

	// Start the server on port 8080
	router.Run(":8080")
}
