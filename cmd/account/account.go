package account

import (
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"github.com/hojin-kr/haru/cmd/trace"
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

// Get Account ID 조회
// @Summary      account ID 조회 (현재 SNS 안하니 미사용)
// @Description  account ID 조회, SNS 연동시 SNS ID로 Account ID 조회
// @Accept       json
// @Tags         account
// @Param        ID   query      int64  true  "ID"
// @Success      200              {array}  account.Account    "ok"
// @Router       /account [get]
func Get(c *gin.Context) {
	tracer := trace.New(os.Stdout)
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
	// ID로 Account 정보 조회
	key := datastore.IDKey("Account", account.ID, nil)
	err = datastoreClient.Get(c.Request.Context(), datastore.IDKey("Account", account.ID, nil), &account)
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not Account get:"+err.Error())
	}
	account.ID = key.ID
	tracer.Trace("Get Account ID ")
	c.JSON(http.StatusOK, account)
}

// Post Create Account ID
// @Summary      account ID 생성
// @Description  account ID 생성
// @Accept       json
// @Tags         account
// @Success      200              {array}  account.Account    "ok"
// @Router       /account [post]
func Post(c *gin.Context) {
	tracer := trace.New(os.Stdout)
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
	account.RegisterTimestamp = time.Now().Unix()
	// Putting an entity into the datastore under an incomplete key will cause a unique key to be generated for that entity, with a non-zero IntID.
	key, err := datastoreClient.Put(c.Request.Context(), datastore.IncompleteKey("Account", nil), &Account{RegisterTimestamp: account.RegisterTimestamp})
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not Account Put:"+err.Error())
		return
	}
	account.ID = key.ID
	tracer.Trace("Create New Account ID ")
	c.JSON(http.StatusOK, account)
}
