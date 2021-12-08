package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

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
	r.POST("/profile/get", getProfile)
	r.POST("/profile/set", setProfile)
	r.Run()
}

// Account account infomation
type Account struct {
	Created  int64
	DeviceID string
	GoogleID string
	AppleID  string
	LineID   string
	KakaoID  string
	ID       int64 `datastore:"-"`
}

// Profile profile inforamtion
type Profile struct {
	Nickname string
	ID       int64 `datastore:"-"`
}

// Wallet wallet
type Wallet struct {
	Coin string
	ID   int64 `datastore:"-"`
}

// getAccount account init
func getAccount(c *gin.Context) {
	var account Account
	if err := c.ShouldBind(&account); err != nil {
		c.String(http.StatusBadRequest, "Should Not Bind:"+err.Error())
		return
	}
	ctx := context.Background()
	if account.ID == -1 {
		account.Created = time.Now().Unix()
		// Putting an entity into the datastore under an incomplete key will cause a unique key to be generated for that entity, with a non-zero IntID.
		key, err := datastoreClient.Put(ctx, datastore.IncompleteKey("Account", nil), &Account{Created: account.Created})
		if err != nil {
			c.String(http.StatusBadRequest, "Should Not Account Put:"+err.Error())
			return
		}
		account.ID = key.ID
	} else {
		key := datastore.IDKey("Account", account.ID, nil)
		err := datastoreClient.Get(ctx, datastore.IDKey("Account", account.ID, nil), &account)
		if err != nil {
			c.String(http.StatusBadRequest, "Should Not Account get:"+err.Error())
		}
		account.ID = key.ID
	}
	c.JSON(http.StatusOK, account)
}

// getProfile profile get
func getProfile(c *gin.Context) {
	var profile Profile
	if err := c.ShouldBind(&profile); err != nil {
		c.String(http.StatusBadRequest, "Should Not Bind:"+err.Error())
		return
	}
	ctx := context.Background()
	key := datastore.IDKey("Profile", profile.ID, nil)
	if err := datastoreClient.Get(ctx, key, &profile); err != nil {
		c.String(http.StatusBadRequest, "Should Not Profile get:"+err.Error())
	}
	c.JSON(http.StatusOK, profile)
}

// setProfile profile get
func setProfile(c *gin.Context) {
	var profile Profile
	if err := c.ShouldBind(&profile); err != nil {
		c.String(http.StatusBadRequest, "Should Not Bind:"+err.Error())
		return
	}
	ctx := context.Background()
	key := datastore.IDKey("Profile", profile.ID, nil)
	_, err := datastoreClient.Put(ctx, key, &profile)
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not Profile Put:"+err.Error())
		return
	}
	c.JSON(http.StatusOK, profile)
}
