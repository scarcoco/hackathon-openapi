package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scarcoco/trajectory-api/routers"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

	r := gin.Default()

	routers.Routers(r)

	r.Run(":3000")
}
