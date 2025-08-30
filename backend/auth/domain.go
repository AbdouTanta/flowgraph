package auth

import "go.mongodb.org/mongo-driver/v2/bson"

// For now, I'm using the domain model in the REST layer too.
// Not worth the complexity of separating them yet.

type User struct {
	Id       bson.ObjectID `bson:"_id,omitempty" json:"id"` // MongoDB uses _id as the primary key
	Email    string        `bson:"email" json:"email"`
	Username string        `bson:"username" json:"username"`
	Password string        `bson:"password" json:"password"`
}
