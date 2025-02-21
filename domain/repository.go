package domain

import "context"

// UserRepository defines methods for user persistence.
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
}

// TaskRepository defines methods for task persistence.
type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	FindByUserID(ctx context.Context, userID string) ([]Task, error)
}
