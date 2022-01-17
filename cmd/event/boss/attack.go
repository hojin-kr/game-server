package boss

import (
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"github.com/hojin-kr/haru/cmd/trace"
)

// Attack attack boss infomation
type Attack struct {
	BossID    string `example:"boss-1"`
	Point     int64  `example:"100"`
	Timestamp int64  `example:"1639056738"`
	ID        int64  `datastore:"-" example:"5373899369873408"`
}

// Get 공격 조회
// @Summary      Account ID로 attack 조회
// @Description  Account ID로 attack 조회
// @Accept       json
// @Tags         attack
// @Param        ID   query      int64  true  "ID"
// @Success      200              {array}  boss.Attack    "ok"
// @Router       /attack [get]
func Get(c *gin.Context) {
	tracer := trace.New(os.Stdout)
	var attack Attack
	if err := c.ShouldBind(&attack); err != nil {
		c.String(http.StatusBadRequest, "Should Not Bind:"+err.Error())
		return
	}
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(c.Request.Context(), os.Getenv("PROJECT_ID"))
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not Datastore New Client"+err.Error())
		return
	}
	// ID로 attack 정보 조회
	key := datastore.IDKey("attack", attack.ID, nil)
	err = datastoreClient.Get(c.Request.Context(), datastore.IDKey("attack", attack.ID, nil), &attack)
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not attack get:"+err.Error())
	}
	attack.ID = key.ID
	tracer.Trace("Get attack ID ")
	c.JSON(http.StatusOK, attack)
}

// Post attack set godoc
// @Summary      attack 등록
// @Description  Account ID로 attack 등록
// @Accept       json
// @Tags         attack
// @Param        ID   query      int64  true  "Account ID"
// @Param        BossId   query      string  true  "boss-1"
// @Param        Point   query      string  true  "100"
// @Success      200              {array}  boss.Attack    "ok"
// @Router       /attack [post]
func Post(c *gin.Context) {
	tracer := trace.New(os.Stdout)
	var attack Attack
	if err := c.ShouldBind(&attack); err != nil {
		c.String(http.StatusBadRequest, "Should Not Bind:"+err.Error())
		return
	}
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(c.Request.Context(), os.Getenv("PROJECT_ID"))
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not Datastore New Client"+err.Error())
		return
	}
	attack.Timestamp = time.Now().Unix()
	key := datastore.IncompleteKey("Attack", nil)
	_, err = datastoreClient.Put(c.Request.Context(), key, &attack)
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not attack Put:"+err.Error())
		return
	}
	tracer.Trace("Put attack")
	c.JSON(http.StatusOK, attack)
}
