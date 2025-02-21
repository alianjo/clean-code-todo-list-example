package domain

// Task represents a to-do item.
type Task struct {
	ID          string `json:"id" bson:"_id"`
	UserID      string `json:"user_id" bson:"user_id"`
	Description string `json:"description" bson:"description"`
	Done        bool   `json:"done" bson:"done"`
}
