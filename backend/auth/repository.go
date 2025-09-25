package auth

import (
	"flowgraph/db"
	"flowgraph/utils"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthRepository struct {
	Db *mongo.Database
}

func NewAuthRepository(db *mongo.Database) *AuthRepository {
	return &AuthRepository{
		Db: db,
	}
}

func (r *AuthRepository) Login(email string, password string) (*User, error) {
	// Check for the user in database
	user, err := db.FindOneDocument[User](r.Db, "users", bson.M{"email": email})
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}
	if user == nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	// Compare the hashed password with the provided password
	err = utils.ComparePasswords(user.Password, password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return user, nil
}

func (r *AuthRepository) Register(email, username, hashedPassword string) (*User, error) {
	user := &User{
		Email:    email,
		Username: username,
		Password: hashedPassword,
	}

	userId, err := db.CreateDocument(r.Db, "users", user)
	if err != nil {
		return nil, err
	}

	user.Id = *userId
	return user, nil
}

func (r *AuthRepository) FindUserById(id string) (*User, error) {
	user, err := db.FindDocumentByID[User](r.Db, "users", id)
	return user, err
}

func (r *AuthRepository) FindUserByEmail(email string) (*User, error) {
	user, err := db.FindOneDocument[User](r.Db, "users", bson.M{"email": email})
	return user, err
}

func (r *AuthRepository) FindUserByUsername(username string) (*User, error) {
	user, err := db.FindOneDocument[User](r.Db, "users", bson.M{"username": username})
	return user, err
}
