package profile

import (
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

// Profile profile inforamtion
type Profile struct {
	Nickname string `example:"myNickname"`
	ID       int64  `datastore:"-" example:"5373899369873408"`
}

// Wallet wallet
type Wallet struct {
	Coin string
	ID   int64 `datastore:"-" example:"5373899369873408"`
}

// getProfile profile get godoc
// @Summary      profile 조회
// @Description  Account ID로 profile 조회
// @Accept       json
// @Tags         profile
// @Param        ID   query      int64  true  "Account ID"
// @Success      200              {array}  profile.Profile    "ok"
// @Router       /profile/get [post]
func Get(c *gin.Context) {
	var profile Profile
	if err := c.ShouldBind(&profile); err != nil {
		c.String(http.StatusBadRequest, "Should Not Bind:"+err.Error())
		return
	}
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(c.Request.Context(), os.Getenv("PROJECT_ID"))
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not Datastore New Client"+err.Error())
		return
	}
	key := datastore.IDKey("Profile", profile.ID, nil)
	if err := datastoreClient.Get(c.Request.Context(), key, &profile); err != nil {
		c.String(http.StatusBadRequest, "Should Not Profile get:"+err.Error())
	}
	c.JSON(http.StatusOK, profile)
}

// getProfile profile set godoc
// @Summary      profile 등록
// @Description  Account ID로 profile 등록
// @Accept       json
// @Tags         profile
// @Param        ID   query      int64  true  "Account ID"
// @Param        Nickname   query      string  true  "Nickname"
// @Success      200              {array}  profile.Profile    "ok"
// @Router       /profile/set [post]
func Set(c *gin.Context) {
	var profile Profile
	if err := c.ShouldBind(&profile); err != nil {
		c.String(http.StatusBadRequest, "Should Not Bind:"+err.Error())
		return
	}
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(c.Request.Context(), os.Getenv("PROJECT_ID"))
	c.String(http.StatusBadRequest, "Should Not Datastore New Client"+err.Error())
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not Datastore New Client"+err.Error())
		return
	}
	key := datastore.IDKey("Profile", profile.ID, nil)
	_, err = datastoreClient.Put(c.Request.Context(), key, &profile)
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not Profile Put:"+err.Error())
		return
	}
	c.JSON(http.StatusOK, profile)
}
