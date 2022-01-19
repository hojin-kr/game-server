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
	Point int64 `example:"100"`
	ID    int64 `datastore:"-" example:"5373899369873408"`
}

// Get Boss point
// @Summary      Boss point
// @Description  Boss point
// @Accept       json
// @Tags         attack
// @Param        ID   query      int64  true  "BossID"
// @Success      200              {array}  boss.Attack    "ok"
// @Router       /boss/attack [get]
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
	key := datastore.IDKey("Attack", attack.ID, nil)
	if err := datastoreClient.Get(c.Request.Context(), key, &attack); err != nil {
		c.String(http.StatusBadRequest, "Should Not Attack get:"+err.Error())
	}
	tracer.Trace("Get attack ID ")
	c.JSON(http.StatusOK, attack)
}

// Put Boss point incr
// @Summary      Boss point incr
// @Description  Boss point incr
// @Accept       json
// @Tags         attack
// @Param        ID   query      int64  true  "BossID"
// @Param        Point   query      int64  true  "100"
// @Success      200              {array}  boss.Attack    "ok"
// @Router       /boss/attack [put]
func Put(c *gin.Context) {
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
	var _attack Attack
	key := datastore.IDKey("Attack", attack.ID, nil)
	if err := datastoreClient.Get(c.Request.Context(), key, &_attack); err != nil {
		c.String(http.StatusBadRequest, "Should Not Attack get:"+err.Error())
	}
	_attack.Point += attack.Point
	_, err = datastoreClient.Put(c.Request.Context(), key, &_attack)
	if err != nil {
		c.String(http.StatusBadRequest, "Should Not Attack Put:"+err.Error())
		return
	}
	tracer.Trace("Put attack")
	c.JSON(http.StatusOK, _attack)
}
