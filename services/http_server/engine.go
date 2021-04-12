package http_server

import (
	"fmt"
	cacheService "github.com/ditoking/website-visitors-stats/services/cache"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type Router struct {
	ginEngine *gin.Engine
	cacheSrv *cacheService.Engine
}

func NewEngine(ginEngine *gin.Engine, cacheSrv *cacheService.Engine) *Router {
	e := new(Router)
	e.ginEngine = ginEngine
	e.cacheSrv = cacheSrv

	return e
}

func (r Router) StartListeningAndServe() error {
	r.defineEndpoints()
	return r.start()
}

func (r Router) start() error {
	httpPort, exists := os.LookupEnv("HTTP_PORT")
	if !exists {
		httpPort = "8080"
	}
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", httpPort),
		Handler:        r.ginEngine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.ListenAndServe()
}