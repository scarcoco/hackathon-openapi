package v1

import (
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	v1 := r.Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "v1 ping pong",
		})
	})
}
