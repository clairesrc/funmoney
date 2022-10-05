package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type transactions struct {
	client storeClient
}

type transactionRecord struct {
	ID        string `bson:"_id",omitempty`
    Comment    string `bson:"comment",omitempty`
    Type    string `bson:"type",omitempty`
    Timestamp string `bson:"timestamp",omitempty`
	Value float64 `bson:"value",omitempty`
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

func (t *transactions) Read(query bson.D) ([]transactionRecord, error) {
	result, err := t.client.findOne(TRANSACTIONS_COLLECTION_NAME, query)
    if err!= nil {
        return nil, fmt.Errorf("Can't find transaction record:\n%w", err)
    }

	return []transactionRecord{
		{
            ID:        result["_id"].(string),
            Comment:    result["comment"].(string),
            Timestamp: result["timestamp"].(string),
            Value:      result["value"].(float64),
	}}, nil
}