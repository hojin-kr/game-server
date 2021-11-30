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
	r.POST("/user/init", userInit)
	r.Run()
}

// User user
type User struct {
	ClientID string
	DeviceID string
	ID       int64 `datastore:"-"`
}

// userInit user get or init
func userInit(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	ctx := context.Background()
	q := datastore.NewQuery("User").Filter("ClientID =", user.ClientID).Limit(1)
	users := make([]*User, 0)
	keys, err := datastoreClient.GetAll(ctx, q, &users)
	if err != nil {
		c.String(http.StatusBadRequest, "user get fail")
	}
	if len(users) < 1 {
		users = append(users, &User{ClientID: user.ClientID, DeviceID: user.DeviceID})
		k := datastore.IncompleteKey("User", nil)
		key, err := datastoreClient.Put(ctx, k, users[0])
		users[0].ID = key.ID
		if err != nil {
			c.String(http.StatusBadRequest, "user init fail")
			return
		}
	} else {
		users[0].ID = keys[0].ID
	}
	c.JSON(http.StatusOK, users[0])
}
