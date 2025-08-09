package flows

import (
	"flowgraph/db"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DbFlowRepository struct {
	Db *mongo.Database
}

func NewFlowsRepository(db *mongo.Database) *DbFlowRepository {
	return &DbFlowRepository{
		Db: db,
	}
}

func (r *DbFlowRepository) FindAllFlows() ([]Flow, error) {
	flows, err := db.FindManyDocuments[Flow](r.Db, "flows", nil)
	return flows, err
}

func (r *DbFlowRepository) FindFlowById(id string) (*Flow, error) {
	flow, err := db.FindDocumentByID[Flow](r.Db, "flows", id)
	return flow, err
}

func (r *DbFlowRepository) InsertFlow(flow *Flow) error {
	_, err := db.CreateDocument(r.Db, "flows", flow)
	return err
}
