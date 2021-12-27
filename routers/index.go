package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	v1 "github.com/scarcoco/trajectory-api/controllers/v1"
	v2 "github.com/scarcoco/trajectory-api/controllers/v2"
)

func Routers(r *gin.Engine) {

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	v1.Setup(r)
	v2.Setup(r)
}
