package logsrepo

import (
	"bail/domain/models"
	"context"
	"errors"
	"time"

	irepo "bail/usecases/core/i_repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Repo struct {
	collection *mongo.Collection
}

// repo implements irepo.log
var _ irepo.SystemLog = &Repo{}

// New creates a new Repository for managing logs with the given MongoDB client, database name, and collection name.
func New(client *mongo.Client, dbName, collectionName string) *Repo {
	collection := client.Database(dbName).Collection(collectionName)
	return &Repo{
		collection: collection,
	}
}

// Save adds a new log if it does not exist, else updates the existing one.
func (r *Repo) Save(systemLog *models.SystemLog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Convert the logmodel.log to logDTO
	logDTO := fromSystemLog(systemLog)

	filter := bson.M{"_id": logDTO.id}
	update := bson.M{
		"$set": logDTO,
	}

	_, err := r.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return errors.New("error savin log")
	}
	return nil
}


