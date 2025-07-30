package service

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Service struct {
	Db *mongo.Database
}
