package flows

import (
	"log"

	"github.com/gin-gonic/gin"
)

type FlowRestController struct {
	flowService *FlowService
}

func NewFlowRestController(flowService *FlowService) *FlowRestController {
	return &FlowRestController{
		flowService: flowService,
	}
}

func (h *FlowRestController) GetAllFlows(c *gin.Context) {
	// Get all flows from database
	flows, err := h.flowService.GetAllFlows()
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

func (h *FlowRestController) GetFlowByID(c *gin.Context) {
	flowID := c.Param("id")

	// Find the flow by ID in the database
	flow, err := h.flowService.GetFlowByID(flowID)

	if err != nil {
		c.JSON(404, gin.H{"error": "Flow not found"})
		return
	}

	c.JSON(200, gin.H{"flow": flow})
}

type CreateFlowDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (h *FlowRestController) CreateFlow(c *gin.Context) {
	var body CreateFlowDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("Error binding create flow data: %v", err)
		c.JSON(400, gin.H{"error": "Invalid flow creation input"})
		return
	}

	// Get user from context
	userId, _ := c.Get("user_id")
	if userId == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	flowName := body.Name
	flowDescription := body.Description
	// Check if the user already has a flow with the same name
	// TODO - Scope by user
	existingFlow, err := h.flowService.GetFlowByName(flowName)
	if err != nil {
		log.Printf("Error checking existing flow: %v", err)
		c.JSON(500, gin.H{"error": "Failed to check existing flows"})
		return
	}
	if existingFlow != nil {
		c.JSON(400, gin.H{"error": "Flow with the same name already exists"})
		return
	}

	// Insert the new flow into the database
	objectId, err := h.flowService.CreateFlow(flowName, flowDescription, userId.(string))
	if err != nil {
		log.Printf("Error creating flow: %v", err)
		c.JSON(500, gin.H{"error": "Failed to create flow"})
		return
	}

	c.JSON(201, gin.H{"message": "Flow created successfully", "id": objectId.Hex()})
}

// func (h *FlowRestController) UpdateFlow(c *gin.Context) {
// 	var flow Flow
// 	if err := c.ShouldBindJSON(&flow); err != nil {
// 		log.Printf("Error binding flow data: %v", err)
// 		c.JSON(400, gin.H{"error": "Invalid flow data"})
// 		return
// 	}

// 	// Insert the new flow into the database
// 	if err := h.flowService.CreateFlow(&flow); err != nil {
// 		log.Printf("Error creating flow: %v", err)
// 		c.JSON(500, gin.H{"error": "Failed to create flow"})
// 		return
// 	}

// 	c.JSON(201, gin.H{"message": "Flow created successfully"})
// }
