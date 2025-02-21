package infrastructure

import (
	"context"

	"github.com/alianjo/clean-code-todo-list-example/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)



// MongoUserRepository implements UserRepository for MongoDB.
type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{collection: db.Collection("users")}
}

func (r *MongoUserRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *MongoUserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// MongoTaskRepository implements TaskRepository for MongoDB.
type MongoTaskRepository struct {
	collection *mongo.Collection
}

func NewMongoTaskRepository(db *mongo.Database) *MongoTaskRepository {
	return &MongoTaskRepository{collection: db.Collection("tasks")}
}

func (r *MongoTaskRepository) Create(ctx context.Context, task *domain.Task) error {
	_, err := r.collection.InsertOne(ctx, task)
	return err
}

func (r *MongoTaskRepository) FindByUserID(ctx context.Context, userID string) ([]domain.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	var tasks []domain.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
