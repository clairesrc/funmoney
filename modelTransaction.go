package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
 *	Model representing transactions
 */
type transactions struct {
	client storeClient
}

/**
 *    Transaction objects for store.
 @TODO: This shouldn't know anything about MongoDB. Figure out a way to use a more generic structure that doesn't use primitive.objectD.
 */
type transactionRecord struct {
    ID      *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
    Comment    string `bson:"comment,omitempty"`
    Type    string `bson:"type,omitempty"`
    Timestamp int `bson:"timestamp,omitempty"`
	Value float64 `bson:"value,omitempty"`
}

/**
 *    Create a new transactions model instance.
 *    @param client        A datastore that implements storeClient.
 *    @return transactions   A transactions model instance.
 */
func newTransactionsModel(client storeClient) *transactions {
    return &transactions{
		client: client,
	}
}

/**
 *    Insert a new transaction record into the database.
 *    @param transaction        The transaction record to insert.
 *	  @return result	   		The result of the operation.
 *    @return errors		   	An error if an error occurred.
 */
func (t *transactions) Create(transaction transactionRecord) (*mongo.InsertOneResult, error) {
	result, err := t.client.insertOne(TRANSACTIONS_COLLECTION_NAME, transaction)
	if err!= nil {
		return nil, fmt.Errorf("Can't add transaction record:\n%w", err)
	}

	return result, nil
}

/**
 *    Find transactions from a filter.
 *    @param query           		The filter.
 *    @param limit           		Likmit number of results.
 *	  @return []TransactionRecord 	Resulting transactions from filter.
 *    @return errors		   		An error if an error occurred.
@TODO: This shouldn't know about MongoDB. Handle the querying and limit option in a more generic way in the storeClient implementer i.e. mongo.go.
 */
func (t *transactions) Read(query bson.D, limit int) ([]transactionRecord, error) {
	var records []transactionRecord
	opts := &options.FindOptions{}
	if limit > 0 {
		opts = options.Find().SetLimit(int64(limit))
	}
	opts = opts.SetSort(bson.D{{Key: "_id", Value: -1}})
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


/**
 *    Get the total balance given transactions from a filter.
 *    @param query           		The filter.
 *	  @return sum				 	The sum figure.
 *    @return errors		   		An error if an error occurred.
 @TODO: This shouldn't know anytihng about MongoDB. Refactor this so it uses a more generic format than bson.M as that should be handled by the storeClient implementation i.e. mongo.go
 */
func (t *transactions) Sum(query bson.D) ([]bson.M, error) {

	result, err := t.client.aggregate(TRANSACTIONS_COLLECTION_NAME, query)
    if err!= nil {
        return nil, fmt.Errorf("Can't find transaction record:\n%w", err)
    }


	return result, nil
}
