package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

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

	transactionsClient := newTransactionsModel(store)

	r := gin.Default()
	r.GET("/transactions", func(c *gin.Context) {
		transactions, err := transactionsClient.Read(bson.D{})
		if err!= nil {
			reportError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	})
	_ = r.Run()
}

func reportError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
} 