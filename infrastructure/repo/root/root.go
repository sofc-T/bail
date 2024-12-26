package rootrepo 


import (
	"context"
	"errors"
	"fmt"

	"bail/domain/models"
	irepo "bail/usecases/core/i_repo"

	
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
func (r *Repo) Save(admin models.Root) error {
	// Convert the Usermodel.User to UserDTO
	UserDTO := fromAdmin(admin)

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