package migration

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Function to create a composite index
func createCompositeIndex(collection *mongo.Collection, fieldNames []string) {
	keys := bson.D{}
	for _, fieldName := range fieldNames {
		keys = append(keys, bson.E{Key: fieldName, Value: 1}) // 1 for ascending order
	}

	indexModel := mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatalf("Error creating composite index: %v", err)
	}

	log.Printf("Composite index created for fields: %v\n", fieldNames)
}


// Function to create an index with ID and name fields
func CreateIndexWithIDAndFName(collection *mongo.Collection) {
	createCompositeIndex(collection, []string{"_id", "name"})
}

// Function to create an index with ID and email fields
func CreateIndexWithIDAndEmail(collection *mongo.Collection) {
	createCompositeIndex(collection, []string{"_id", "email"})
}

