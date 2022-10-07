package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Db name in data store
const DBNAME="funmoney"

// Transactions collection name in data store
const TRANSACTIONS_COLLECTION_NAME="transactions"

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

	store, err := newMongoClient("mongodb://root:example@mongo:27017/?maxPoolSize=20&w=majority", DBNAME)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = store.close()
	}()

	transactionsClient := newTransactionsModel(store)

	r := gin.Default()
	
	corsconfig := cors.DefaultConfig()
	corsconfig.AllowAllOrigins = true
	r.Use(cors.New(corsconfig))
	startTimestamp := time.Now().Unix()

	r.GET("/app", func(c *gin.Context) {
		appValues := map[string]interface{}{
			"appName": "FunMoney",
			"currency": currency,
			"cap": cap,
			"lastRestart": startTimestamp,
		}
		c.JSON(http.StatusOK, gin.H{"app": appValues})
	})

	r.GET("/transactions", func(c *gin.Context) {
		transactions, err := transactionsClient.Read(bson.D{}, 12)
		if err!= nil {
			reportError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	})

	r.GET("/balance", func(c *gin.Context) {
		sum, err := transactionsClient.Sum(bson.D{{Key: "value", Value: "$value"}})
		if err!= nil {
			reportError(c, err)
			return
		}

		if len(sum) > 0 {
			c.JSON(http.StatusOK, gin.H{"balance": sum[0]["value"]})
			return
		}

		c.JSON(http.StatusOK, gin.H{"balance": 0})
	})

    r.POST("/transaction", func(c *gin.Context) {
        var transaction transactionRecord
        err := c.BindJSON(&transaction)
		if err!= nil {
            reportError(c, err)
            return
		}

		transaction.Timestamp = int(time.Now().Unix())

		result, err := transactionsClient.Create(transaction)
		if err!= nil {
            reportError(c, err)
            return
		}

		transactionID := result.InsertedID.(primitive.ObjectID)
		transaction.ID = &transactionID

        c.JSON(200, gin.H{"transaction": transaction}) // Your custom response here
    })

	_ = r.Run()
}

func reportError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
