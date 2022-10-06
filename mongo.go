package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoClient struct {
	client *mongo.Client
	database string
}

func newMongoClient(uri, database string) (*mongoClient, error) {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	return &mongoClient{client, database}, nil
}

func (c *mongoClient) close() error {
	if err := c.client.Disconnect(context.TODO()); err != nil {
		return err
	}

	return nil
}


func (c *mongoClient) insertOne(collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	db := c.client.Database(c.database)
	collection := db.Collection(collectionName)
    result, err := collection.InsertOne(context.TODO(), document)
    if err!= nil {
		return nil, fmt.Errorf("Error running insertOne for document:\n%w", err)
	}

	return result, nil
}

func (c *mongoClient) insert(collectionName string, documents []interface{}) (*mongo.InsertManyResult, error) {
	db := c.client.Database(c.database)
	collection := db.Collection(collectionName)
    result, err := collection.InsertMany(context.TODO(), documents)
    if err!= nil {
        return nil, fmt.Errorf("Error running insertMany for documents:\n%w", err)
	}

	return result, nil
}

func (c *mongoClient) update(collectionName string, update, filter bson.D) (*mongo.UpdateResult, error) {
	db := c.client.Database(c.database)
	collection := db.Collection(collectionName)
    result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err!= nil {
		return nil, fmt.Errorf("Error running updateOne:\n%w", err)
	}

	return result, nil
}

func (c *mongoClient) delete(collectionName string, filter bson.D) (*mongo.DeleteResult, error) {
	db := c.client.Database(c.database)
	collection := db.Collection(collectionName)
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("Error deleting:\n%w", err)
	}

	return result, nil
}

func (c *mongoClient) find(collectionName string, query bson.D, opts *options.FindOptions) (*mongo.Cursor, error) {
	db := c.client.Database(c.database)
	collection := db.Collection(collectionName)

	cursor, err := collection.Find(context.TODO(), query, opts)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("Error running find for query %s:\n%w\n", query, err)
	}

	return cursor, nil
}

func (c *mongoClient) aggregate(collectionName string, query bson.D) ([]bson.M, error) {
	db := c.client.Database(c.database)
	collection := db.Collection(collectionName)
	matchStage := bson.D{{Key: "$project", Value: query}}
	groupStage := bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: ""},
        {Key: "value", Value: bson.D{{Key: "$sum", Value: "$value"}}},
        }}}

	result, err := collection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		return nil, fmt.Errorf("Error running aggregate for pipeline %s:\n%w\n", query, err)
	}

	var sum []bson.M
	if err = result.All(context.TODO(), &sum); err != nil {
		panic(err)
	}

	return sum, nil
}
