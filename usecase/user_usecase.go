package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/alianjo/clean-code-todo-list-example/domain"
)

// UserUsecase encapsulates user-related business logic.
type UserUsecase struct {
	userRepo domain.UserRepository
}

// NewUserUsecase creates a new UserUsecase.
func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

// CreateUser adds a new user.
func (u *UserUsecase) CreateUser(ctx context.Context, name, email string) (*domain.User, error) {
	user := &domain.User{
		ID:    generateID(), // Assume a simple ID generator for now.
		Name:  name,
		Email: email,
	}
	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

// Helper function (in a real project, use UUID or similar).
func generateID() string {
	return "user-" + fmt.Sprintf("%d", time.Now().UnixNano())
}
