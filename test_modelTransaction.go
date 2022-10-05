package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

const DBNAME = "funmoney"
const TRANSACTIONS_COLLECTION_NAME="transactions"
const APP_COLLECTION_NAME="app"

func test_modelTransactions() {
	var cap = os.Getenv("CAP")
	if cap == "" {
      cap = "100"
	}
	var currency = os.Getenv("CURRENCY")
	if currency == "" {
		currency="USD"
	}

	err := godotenv.Load(".config.env")
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	log.Printf("Using %s as currency", currency)
	log.Printf("Using %s as cap", cap)

	store, err := newMongoClient("mongodb://root:example@mongo:27017/?maxPoolSize=20&w=majority", DBNAME)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = store.close()
	}()

	transactions := newTransactionsModel(store)

	// Add sample record
	log.Printf("Adding sample record")
	insertedID, err := transactions.Create(transactionRecord{
		Value:    100,
		Type:  "credit",
		Comment: "Sample comment",
		Timestamp: fmt.Sprint((time.Now().Unix())),
	})
	if err!= nil {
		panic(fmt.Errorf("Can't add sample record:\n%w", err))
	}

	fmt.Println(insertedID)


	records, err := transactions.Read(bson.D{})
	if err != nil {
		panic("Can't get initial records")
	}
	fmt.Println(records)

	sum, err := transactions.Sum(bson.D{{"value", "$value"}})
	if err != nil {
		panic(fmt.Errorf("Can't get sum: %w", err))
	}
	fmt.Println(sum)

}