package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcoco/trajectory-api/controllers/v1/reporter"
)

func Setup(r *gin.Engine) {

	v1 := r.Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "v1 ping pong",
		})
	})
	v1.GET("/reporter/list", func(c *gin.Context) {
		c.JSON(200, reporter.List())
	})
	v1.GET("/reporter/:id", func(c *gin.Context) {
		c.JSON(200, reporter.Detail())
	})
	v1.GET("/reporter/:id/location", func(c *gin.Context) {
		c.JSON(200, reporter.Location())
	})
}
