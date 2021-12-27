package main

import (
	"github.com/gin-gonic/gin"
	"github.com/scarcoco/trajectory-api/routers"
)

func main() {
	r := gin.Default()

	routers.Routers(r)

	r.Run(":3001")
}
