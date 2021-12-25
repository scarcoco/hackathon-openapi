package main

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcoco/hackathon-openapi/routers"
)

func main() {
	r := gin.Default()

	routers.Routers(r)

	r.Run(":3000")
}
