package userrepo

import (
	"context"
	"errors"
	"fmt"

	"bail/domain/models"
	irepo "bail/usecases/core/i_repo"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repo defines the MongoDB repository for Users.
type Repo struct {
	collection *mongo.Collection
}

// repo implements irepo.User
var _ irepo.User = &Repo{}

// New creates a new Repository for managing Users with the given MongoDB client, database name, and collection name.
func New(client *mongo.Client, dbName, collectionName string) *Repo {
	collection := client.Database(dbName).Collection(collectionName)
	return &Repo{
		collection: collection,
	}
}

// Save adds a new User if it does not exist, else updates the existing one.
func (r *Repo) Save(User *models.User) error {
	// Convert the Usermodel.User to UserDTO
	UserDTO := FromUser(User)

	filter := bson.M{"_id": UserDTO.Id}
	update := bson.M{
		"$set": UserDTO,
	}

	_, err := r.collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("error saving User: %w", err)
	}
	return nil
}

// FindByID retrieves a User by its ID.
func (r *Repo) FindById(id uuid.UUID) (*models.User, error) {
	filter := bson.M{"_id": id}

	var UserDTO userDTO
	err := r.collection.FindOne(context.Background(), filter).Decode(&UserDTO)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("User not found")
		}
		return nil, fmt.Errorf("error finding User by ID: %w", err)
	}

	User := ToUser(&UserDTO)
	return User, nil
}

// FindByEmail retrieves a User by its email.
func (r *Repo) FindByEmail(email string) (*models.User, error) {
	filter := bson.M{"email": email}

	var UserDTO userDTO
	err := r.collection.FindOne(context.Background(), filter).Decode(&UserDTO)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("User not found")
		}
		return nil, fmt.Errorf("error finding User by email: %w", err)
	}

	User := ToUser(&UserDTO)
	return User, nil
}

//delete by id
func (r *Repo) Delete(id uuid.UUID) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting User by ID: %w", err)
	}
	return nil
}

//get list of employees
func (r *Repo) GetAll(page int) ([]*models.User, error) {
	var users []*models.User
	findOptions := options.Find()
	findOptions.SetSkip(int64(page))
	findOptions.SetLimit(30)
	cursor, err := r.collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error getting list of Users: %w", err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var UserDTO userDTO
		err := cursor.Decode(&UserDTO)
		if err != nil {
			return nil, fmt.Errorf("error decoding User: %w", err)
		}
		User := ToUser(&UserDTO)
		users = append(users, User)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over Users: %w", err)
	}
	return users, nil
}