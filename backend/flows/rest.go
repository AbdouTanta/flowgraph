package flows

import (
	"flowgraph/utils"

	"github.com/gin-gonic/gin"
)

type FlowRestController struct {
	flowService *FlowService
	responder   *utils.HTTPResponder
}

func NewFlowRestController(flowService *FlowService, responder *utils.HTTPResponder) *FlowRestController {
	return &FlowRestController{
		flowService: flowService,
		responder:   responder,
	}
}

func (h *FlowRestController) GetAllFlows(c *gin.Context) {
	// Get all flows from database
	flows, err := h.flowService.GetAllFlows()
	if err != nil {
		h.responder.InternalError(c, "Failed to retrieve flows", err)
		return
	}

	// Return the flows as JSON
	h.responder.SuccessWithMessage(c, 200, "Flows retrieved successfully", gin.H{"flows": flows})
}

func (h *FlowRestController) GetFlowByID(c *gin.Context) {
	flowID := c.Param("id")

	flow, err := h.flowService.GetFlowByID(flowID)
	if err != nil {
		h.responder.InternalError(c, "Failed to retrieve flow", err)
		return
	}

	if flow == nil {
		h.responder.NotFound(c, "Flow not found")
		return
	}

	h.responder.SuccessWithMessage(c, 200, "Flow retrieved successfully", gin.H{"flow": flow})
}

type CreateFlowDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (h *FlowRestController) CreateFlow(c *gin.Context) {
	var body CreateFlowDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		h.responder.BadRequest(c, "Invalid flow creation input", err)
		return
	}

	// Get user from context
	userId, _ := c.Get("user_id")
	if userId == nil {
		h.responder.Unauthorized(c, "Unauthorized")
		return
	}

	flowName := body.Name
	flowDescription := body.Description
	// Check if the user already has a flow with the same name
	// TODO - Scope by user
	existingFlow, err := h.flowService.GetFlowByName(flowName)
	if err != nil {
		h.responder.InternalError(c, "Failed to check existing flows", err)
		return
	}
	if existingFlow != nil {
		h.responder.BadRequest(c, "Flow with the same name already exists", nil)
		return
	}

	// Insert the new flow into the database
	objectId, err := h.flowService.CreateFlow(flowName, flowDescription, userId.(string))
	if err != nil {
		h.responder.InternalError(c, "Failed to create flow", err)
		return
	}

	h.responder.SuccessWithMessage(c, 201, "Flow created successfully", gin.H{"id": objectId.Hex()})
}
