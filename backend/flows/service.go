package flows

type FlowService struct {
	flowRepository FlowRepository
}

func NewFlowService(flowRepository FlowRepository) *FlowService {
	return &FlowService{
		flowRepository: flowRepository,
	}
}

func (s *FlowService) GetAllFlows() ([]Flow, error) {
	flows, err := s.flowRepository.FindAllFlows()
	return flows, err
}

func (s *FlowService) GetFlowByID(id string) (*Flow, error) {
	flow, err := s.flowRepository.FindFlowById(id)
	return flow, err
}

func (s *FlowService) CreateFlow(flow *Flow) error {
	err := s.flowRepository.InsertFlow(flow)
	return err
}
