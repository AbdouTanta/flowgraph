package flows

import (
	"flowgraph/db"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FlowRepository struct {
	Db *mongo.Database
}

func NewFlowsRepository(db *mongo.Database) *FlowRepository {
	return &FlowRepository{
		Db: db,
	}
}

func (r *FlowRepository) FindAllFlows() ([]Flow, error) {
	flows, err := db.FindManyDocuments[Flow](r.Db, "flows", nil)
	return flows, err
}

func (r *FlowRepository) FindFlowById(id string) (*Flow, error) {
	flow, err := db.FindDocumentByID[Flow](r.Db, "flows", id)
	return flow, err
}

func (r *FlowRepository) FindFlowByName(name string) (*Flow, error) {
	flow, err := db.FindOneDocument[Flow](r.Db, "flows", bson.M{"name": name})
	return flow, err
}

func (r *FlowRepository) InsertFlow(flow *Flow) (*bson.ObjectID, error) {
	objectId, err := db.CreateDocument(r.Db, "flows", flow)
	return objectId, err
}
