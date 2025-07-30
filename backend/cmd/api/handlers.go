package main

import (
	"flowgraph/internal/db"
	"log"

	"github.com/gin-gonic/gin"
)

func (app *application) GetAllFlows(c *gin.Context) {
	// Get all flows from database
	flows, err := db.FindManyDocuments[db.Flow](app.service.Db, "flows", nil)
	if err != nil {
		log.Printf("Error retrieving flows: %v", err)
		c.JSON(500, gin.H{
			"error": "Failed to retrieve flows",
		})
		return
	}

	// Return the flows as JSON
	c.JSON(200, gin.H{"flows": flows})
}

func (app *application) GetFlowByID(c *gin.Context) {
	flowID := c.Param("id")

	// Find the flow by ID in the database
	flow, err := db.FindDocumentByID[db.Flow](app.service.Db, "flows", flowID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Flow not found"})
		return
	}

	c.JSON(200, gin.H{"flow": flow})
}

func (app *application) CreateFlow(c *gin.Context) {
	var flow db.Flow
	if err := c.ShouldBindJSON(&flow); err != nil {
		log.Printf("Error binding flow data: %v", err)
		c.JSON(400, gin.H{"error": "Invalid flow data"})
		return
	}

	// Insert the new flow into the database
	if _, err := db.CreateDocument(app.service.Db, "flows", flow); err != nil {
		log.Printf("Error creating flow: %v", err)
		c.JSON(500, gin.H{"error": "Failed to create flow"})
		return
	}

	c.JSON(201, gin.H{"message": "Flow created successfully"})
}
