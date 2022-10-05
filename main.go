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

func main() {
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

	var record map[string]interface{}

	// Check if first run
	record, err = store.findOne(APP_COLLECTION_NAME, bson.D{})
	if err!= nil {
        panic(fmt.Errorf("Can't check for app record:\n%w", err))
    }
	if len(record) == 0 {
		// Add sample record
		log.Printf("Adding sample record")
        if _, err := store.insertOne(TRANSACTIONS_COLLECTION_NAME, bson.D{
            {Key: "comment", Value: "sample"},
			{Key: "type", Value: "credit"},
			{Key: "value", Value: 100},
			{Key: "timestamp", Value: time.Now().Unix()},
		}); err!= nil {
			panic(fmt.Errorf("Can't add sample record:\n%w", err))
		}
		record, err = store.findOne(TRANSACTIONS_COLLECTION_NAME, bson.D{
			{Key: "comment", Value: "sample"},
		})
		if err != nil {
			panic("Can't get sample record after inserting")
		}

		// Add firstRun to app collection
		log.Printf("Adding firstRun record")
        _, err := store.insertOne(APP_COLLECTION_NAME, bson.D{
            {Key: "firstRun", Value: true},
		}) 
		if err != nil {
			panic(fmt.Errorf("Can't add firstRun record:\n%w", err))
		}
    }

	fmt.Println(record)

}