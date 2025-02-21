package interfaces

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/alianjo/clean-code-todo-list-example/usecase"
	"github.com/gorilla/mux"
)

// Handler holds dependencies for HTTP routes.
type Handler struct {
	userUsecase *usecase.UserUsecase
	taskUsecase *usecase.TaskUsecase
}

// NewHandler initializes the handler.
func NewHandler(userUC *usecase.UserUsecase, taskUC *usecase.TaskUsecase) *Handler {
	return &Handler{userUsecase: userUC, taskUsecase: taskUC}
}

// CreateUser handles POST /users.
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	user, err := h.userUsecase.CreateUser(context.Background(), input.Name, input.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// CreateTask handles POST /tasks.
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID      string `json:"user_id"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	task, err := h.taskUsecase.CreateTask(context.Background(), input.UserID, input.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

// GetTasks handles GET /tasks/:userID.
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]
	tasks, err := h.taskUsecase.GetTasks(context.Background(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}
