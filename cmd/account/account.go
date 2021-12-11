package account

import (
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

// Account account infomation
type Account struct {
	RegisterTimestamp int64 `example:"1639056738"`
	DeviceID          string
	GoogleID          string
	AppleID           string
	LineID            string
	KakaoID           string
	ID                int64 `datastore:"-" example:"5373899369873408"`
}

// Get godoc
// @Summary      account 조회 및 생성
// @Description  account 조회 및 생성 ID -1 로 요청시 account 생성 후 ID 반환
// @Accept       json
// @Tags         account
// @Param        ID   query      int64  true  "ID"
// @Success      200              {array}  account.Account    "ok"
// @Router       /account/get [post]
func Get(c *gin.Context) {
	var account Account
	if err := c.ShouldBind(&account); err != nil {
		c.String(http.StatusBadRequest, "Should Not Bind:"+err.Error())
		return
	}
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(c.Request.Context(), os.Getenv("PROJECT_ID"))
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not Datastore New Client"+err.Error())
		return
	}
	if account.ID == -1 {
		account.RegisterTimestamp = time.Now().Unix()
		// Putting an entity into the datastore under an incomplete key will cause a unique key to be generated for that entity, with a non-zero IntID.
		key, err := datastoreClient.Put(c.Request.Context(), datastore.IncompleteKey("Account", nil), &Account{RegisterTimestamp: account.RegisterTimestamp})
		if err != nil {
			c.String(http.StatusBadRequest, "Should Not Account Put:"+err.Error())
			return
		}
		account.ID = key.ID
	} else {
		key := datastore.IDKey("Account", account.ID, nil)
		err := datastoreClient.Get(c.Request.Context(), datastore.IDKey("Account", account.ID, nil), &account)
		if err != nil {
			c.String(http.StatusBadRequest, "Should Not Account get:"+err.Error())
		}
		account.ID = key.ID
	}
	c.JSON(http.StatusOK, account)
}
