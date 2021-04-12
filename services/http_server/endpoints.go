package http_server

import (
	"github.com/gin-gonic/gin"
	"time"
)

func (r Router) defineEndpoints() {
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.ginEngine.Use(gin.Recovery())

	// Add requests a timeout of 10 seconds prevents from blocking
	r.ginEngine.Use(timeoutMiddleware(time.Second * 10))

	// Health Check
	r.ginEngine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Version 1 endpoints
	v1 := r.ginEngine.Group("v1/website-visitor-stats")

	// Endpoints
	v1.POST("/new-interaction", r.HandleNewWebsiteInteraction)
	v1.POST("/new-interaction/", r.HandleNewWebsiteInteraction)
	v1.GET("/page-unique-visitors-count", r.HandleGetPageUniqueVisitorsCount)
	v1.GET("/page-unique-visitors-count/", r.HandleGetPageUniqueVisitorsCount)
}
