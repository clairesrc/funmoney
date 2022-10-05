package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type storeClient interface {
	close() error
	findOne(collectionName string, query bson.D) (map[string]interface{}, error)
	find(collectionName string, query bson.D) ([]interface{}, error)
	insertOne(collectionName string, document interface{}) (*mongo.InsertOneResult, error)
	insert(collectionName string, documents []interface{}) (*mongo.InsertManyResult, error)
	update(collectionName string, update, filter bson.D) (*mongo.UpdateResult, error)
	delete(collectionName string, filter bson.D) (*mongo.DeleteResult, error)
}