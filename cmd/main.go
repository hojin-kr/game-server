package main

import (
	"os"

	"github.com/gin-gonic/gin"
	pkgAccount "github.com/hojin-kr/haru/cmd/account"
	docs "github.com/hojin-kr/haru/cmd/docs"
	pkgEvent "github.com/hojin-kr/haru/cmd/event/boss"
	pkgProfile "github.com/hojin-kr/haru/cmd/profile"
	"github.com/hojin-kr/haru/cmd/trace"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Game Server Basic API
// @version         1.0
// @description     This is a game server basic.

// @contact.name   Hojin Jang
// @contact.url    https://github.com/hojin-kr
// @contact.email  jhj377@gmail.com
func main() {
	// Set this in app.yaml when running in production.
	r := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.BasePath = "/v1"
	v1 := r.Group("/v1")
	{
		v1.GET("account", pkgAccount.Get)
		v1.GET("profile", pkgProfile.Get)
		v1.POST("profile", pkgProfile.Post)
		v1.GET("boss/attack", pkgEvent.Get)
		v1.POST("boss/attack", pkgEvent.Post)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		tracer := trace.New(os.Stdout)
		tracer.Trace("Ping")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
