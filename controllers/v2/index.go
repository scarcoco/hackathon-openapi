package v2

import (
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	v2 := r.Group("/v2")

	v2.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "v2 ping pong",
		})
	})
}
