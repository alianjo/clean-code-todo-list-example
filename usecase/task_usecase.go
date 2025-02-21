package usecase

import (
	"context"

	"github.com/alianjo/clean-code-todo-list-example/domain"
)

// TaskUsecase encapsulates task-related business logic.
type TaskUsecase struct {
	taskRepo domain.TaskRepository
	userRepo domain.UserRepository // To validate user existence.
}

// NewTaskUsecase creates a new TaskUsecase.
func NewTaskUsecase(taskRepo domain.TaskRepository, userRepo domain.UserRepository) *TaskUsecase {
	return &TaskUsecase{taskRepo: taskRepo, userRepo: userRepo}
}

// CreateTask adds a new task for a user.
func (t *TaskUsecase) CreateTask(ctx context.Context, userID, description string) (*domain.Task, error) {
	// Validate user exists.
	if _, err := t.userRepo.FindByID(ctx, userID); err != nil {
		return nil, err
	}
	task := &domain.Task{
		ID:          generateID(),
		UserID:      userID,
		Description: description,
		Done:        false,
	}
	if err := t.taskRepo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

// GetTasks retrieves tasks for a user.
func (t *TaskUsecase) GetTasks(ctx context.Context, userID string) ([]domain.Task, error) {
	return t.taskRepo.FindByUserID(ctx, userID)
}
