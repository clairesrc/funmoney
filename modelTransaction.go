package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type transactions struct {
	client storeClient
}

type transactionRecord struct {
    ID      *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
    Comment    string `bson:"comment,omitempty"`
    Type    string `bson:"type,omitempty"`
    Timestamp string `bson:"timestamp,omitempty"`
	Value float64 `bson:"value,omitempty"`
}

func newTransactionsModel(client storeClient) *transactions {
    return &transactions{
		client: client,
	}
}

func (t *transactions) Create(transaction transactionRecord) (interface{}, error) {
	result, err := t.client.insertOne(TRANSACTIONS_COLLECTION_NAME, transaction)
	if err!= nil {
		return nil, fmt.Errorf("Can't add transaction record:\n%w", err)
	}
	
	return result.InsertedID, nil
}

func (t *transactions) Read(query bson.D, limit int) ([]transactionRecord, error) {
	var records []transactionRecord
	opts := &options.FindOptions{}
	if limit > 0 {
		opts = options.Find().SetLimit(int64(limit))
	}
	result, err := t.client.find(TRANSACTIONS_COLLECTION_NAME, query, opts)
    if err!= nil {
        return nil, fmt.Errorf("Can't find transaction record:\n%w", err)
    }

	for result.Next(context.TODO()) {
		record := transactionRecord{}
		err := result.Decode(&record)
		if err != nil {
			return nil, fmt.Errorf("Can't decode transaction record: %w", err)
		}
		records = append(records, record)
	}

	return records, nil
}

func (t *transactions) Sum(query bson.D) ([]bson.M, error) {

	result, err := t.client.aggregate(TRANSACTIONS_COLLECTION_NAME, query)
    if err!= nil {
        return nil, fmt.Errorf("Can't find transaction record:\n%w", err)
    }

	return result, nil
}
