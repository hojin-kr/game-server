package boss

import (
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"github.com/hojin-kr/haru/cmd/trace"
)

// Attack attack boss infomation
type Attack struct {
	BossID int64 `example:"1"`
	Point  int64 `example:"100"`
	ID     int64 `datastore:"-" example:"5373899369873408"`
}

// Get Boss가 공격 당한 정보 조회
// @Summary      Boss 공격 당한 정보 조회
// @Description  Boss 공격 당한 정보 조회, 공격당한 전체 총합 point
// @Accept       json
// @Tags         attack
// @Param        BossID   query      int64  true  "BossID"
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
	var attacks []Attack
	q := datastore.NewQuery("Attack").Filter("BossID =", attack.BossID)
	if _, err := datastoreClient.GetAll(c.Request.Context(), q, &attacks); err != nil {
		c.String(http.StatusBadRequest, "Should Not attack getAll:"+err.Error())
	}
	for _, v := range attacks {
		attack.Point += v.Point
	}
	tracer.Trace("Get attack ID ")
	c.JSON(http.StatusOK, attack)
}

// Post attack 등록
// @Summary      attack 등록
// @Description  attack 등록
// @Accept       json
// @Tags         attack
// @Param        ID   query      int64  true  "Account ID"
// @Param        BossId   query      int64  true  "1"
// @Param        Point   query      int64  true  "100"
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
	key := datastore.IncompleteKey("Attack", nil)
	_, err = datastoreClient.Put(c.Request.Context(), key, &attack)
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not attack Put:"+err.Error())
		return
	}
	tracer.Trace("Put attack")
	c.JSON(http.StatusOK, attack)
}
