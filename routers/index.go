package routers

import (
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

	v1.Setup(r)
	v2.Setup(r)
}
