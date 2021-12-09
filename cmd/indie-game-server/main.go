package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	docs "github.com/hojin-kr/indie-game-server-architecture/cmd/indie-game-server/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var datastoreClient *datastore.Client

// @title           Game Server Basic API
// @version         1.0
// @description     This is a game server basic.

// @contact.name   Hojin Jang
// @contact.url    https://github.com/hojin-kr
// @contact.email  jhj377@gmail.com
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
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.POST("get", getAccount)
		}
		profiles := v1.Group("/profiles")
		{
			profiles.POST("get", getProfile)
			profiles.POST("set", setProfile)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

// Account account infomation
type Account struct {
	RegisterTimestamp int64
	DeviceID          string
	GoogleID          string
	AppleID           string
	LineID            string
	KakaoID           string
	ID                int64 `datastore:"-"`
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

// getAccount godoc
// @Summary      get an account
// @Description  get or Register account by ID, ID에 -1을 넘기면 계정 생성
// @Accept       json
// @Tags         accounts
// @Param        ID   query      int64  true  "Account ID"
// @Success      200              {Account}  model.Account    "ok"
// @Router       /accounts/get [post]
func getAccount(c *gin.Context) {
	var account Account
	if err := c.ShouldBind(&account); err != nil {
		c.String(http.StatusBadRequest, "Should Not Bind:"+err.Error())
		return
	}
	ctx := context.Background()
	if account.ID == -1 {
		account.RegisterTimestamp = time.Now().Unix()
		// Putting an entity into the datastore under an incomplete key will cause a unique key to be generated for that entity, with a non-zero IntID.
		key, err := datastoreClient.Put(ctx, datastore.IncompleteKey("Account", nil), &Account{RegisterTimestamp: account.RegisterTimestamp})
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

// getProfile profile get godoc
// @Summary      get an profile
// @Description  get profile by ID
// @Accept       json
// @Tags         profiles
// @Param        ID   query      int64  true  "Account ID"
// @Success      200              {Profile}  model.Profile    "ok"
// @Router       /profiles/get [post]
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

// getProfile profile set godoc
// @Summary      set an profile
// @Description  set profile by ID
// @Accept       json
// @Tags         profiles
// @Param        ID   query      int64  true  "Account ID"
// @Param        Nickname   query      string  true  "Nickname"
// @Success      200              {Profile}  model.Profile    "ok"
// @Router       /profiles/set [post]
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
