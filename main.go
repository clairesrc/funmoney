package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

const DBNAME="funmoney"
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

		c.JSON(http.StatusOK, gin.H{"balance": sum[0]["value"]})
	})

	_ = r.Run()
}

func reportError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
