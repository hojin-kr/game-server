package main

import (
	"os"

	"github.com/gin-gonic/gin"
	pkgAccount "github.com/hojin-kr/game-server/cmd/account"
	docs "github.com/hojin-kr/game-server/cmd/docs"
	pkgProfile "github.com/hojin-kr/game-server/cmd/profile"
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
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/v1")
	{
		v1.GET("account", pkgAccount.Get)
		v1.GET("profile", pkgProfile.Get)
		v1.POST("profile", pkgProfile.Post)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
