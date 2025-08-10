package auth

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DbAuthRepository struct {
	Db *mongo.Database
}

func NewAuthRepository(db *mongo.Database) *DbAuthRepository {
	return &DbAuthRepository{
		Db: db,
	}
}

func (r *DbAuthRepository) Login(user *User) (*User, error) {
	// For now, let's simulate a login by checking hardcoded credentials
	if user.Email != "test@test.com" || user.Password != "test" {
		return nil, fmt.Errorf("invalide email or password")
	}

	return user, nil
}
