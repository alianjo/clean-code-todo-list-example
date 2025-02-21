package domain

// User represents a user entity.
type User struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
