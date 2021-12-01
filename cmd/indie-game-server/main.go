package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

var datastoreClient *datastore.Client

func main() {
	ctx := context.Background()
	// Set this in app.yaml when running in production.
	projectID := os.Getenv("PROJECT_ID")
	var err error
	datastoreClient, err = datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/account/get", getAccount)
	r.Run()
}

// Account account infomation
type Account struct {
	ClientID string
	DeviceID string
	ID       string `datastore:"-"`
}

// Profile profile inforamtion
type Profile struct {
	Nickname string
}

// Wallet wallet
type Wallet struct {
	Coin string
}

// getAccount account get or init
func getAccount(c *gin.Context) {
	var account Account
	if err := c.ShouldBind(&account); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	ctx := context.Background()
	q := datastore.NewQuery("Account").Filter("ClientID =", account.ClientID).Limit(1)
	accounts := make([]*Account, 0)
	keys, err := datastoreClient.GetAll(ctx, q, &accounts)
	if err != nil {
		c.String(http.StatusBadRequest, "Account get fail")
	}
	if len(accounts) < 1 {
		accounts = append(accounts, &Account{ClientID: account.ClientID, DeviceID: account.DeviceID})
		k := datastore.IncompleteKey("Account", nil)
		key, err := datastoreClient.Put(ctx, k, accounts[0])
		accounts[0].ID = key.String()
		if err != nil {
			c.String(http.StatusBadRequest, "Account init fail")
			return
		}
	} else {
		accounts[0].ID = keys[0].String()
	}
	c.JSON(http.StatusOK, accounts[0])
}
