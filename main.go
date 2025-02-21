package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/alianjo/clean-code-todo-list-example/infrastructure"
	"github.com/alianjo/clean-code-todo-list-example/interfaces"
	"github.com/alianjo/clean-code-todo-list-example/usecase"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Get MongoDB URI from environment variable, with fallback for local dev
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" // Default for local MongoDB
	}

	// MongoDB setup
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("todo")

	// Dependency injection
	userRepo := infrastructure.NewMongoUserRepository(db)
	taskRepo := infrastructure.NewMongoTaskRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	taskUC := usecase.NewTaskUsecase(taskRepo, userRepo)
	handler := interfaces.NewHandler(userUC, taskUC)

	// Router setup
	r := mux.NewRouter()
	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{userID}", handler.GetTasks).Methods("GET")

	// Start server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
