package flows

import "go.mongodb.org/mongo-driver/v2/bson"

type FlowService struct {
	flowRepository FlowRepository
}

func NewFlowService(flowRepository *FlowRepository) *FlowService {
	return &FlowService{
		flowRepository: *flowRepository,
	}
}

func (s *FlowService) GetAllFlows() ([]Flow, error) {
	// Initialize empty slice to avoid returning nil in handler
	flows := make([]Flow, 0)
	results, err := s.flowRepository.FindAllFlows()
	if results == nil {
		results = flows
	}
	return results, err
}

func (s *FlowService) GetFlowByID(id string) (*Flow, error) {
	flow, err := s.flowRepository.FindFlowById(id)
	return flow, err
}

func (s *FlowService) GetFlowByName(name string) (*Flow, error) {
	flow, err := s.flowRepository.FindFlowByName(name)
	return flow, err
}

func (s *FlowService) CreateFlow(name string, description string, userId string) (*bson.ObjectID, error) {
	flow := &Flow{
		Name:        name,
		Description: description,
		UserId:      userId,
	}
	objectId, err := s.flowRepository.InsertFlow(flow)
	return objectId, err
}
