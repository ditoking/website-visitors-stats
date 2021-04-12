package main

import (
	cacheService "github.com/ditoking/website-visitors-stats/services/cache"
	"github.com/ditoking/website-visitors-stats/services/http_server"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

var (
	ginEngine    *gin.Engine
	httpRouter   *http_server.Router
	serviceCache *cache.Cache
	cacheSrv     *cacheService.Engine
)

func init() {
	initCache()
	initGin()

	cacheSrv = cacheService.NewEngine(serviceCache)
	httpRouter = http_server.NewEngine(ginEngine, cacheSrv)

}

func main() {
	if err := httpRouter.StartListeningAndServe(); err != nil {
		log.Fatalf("error initializing router: %s", err.Error())
	}
}

// --- Inits --- //
func initGin() {
	// Init Gin Router
	ginEngine = gin.New()
}

func initCache() {
	// Init in memory cache
	serviceCache = cache.New(cache.NoExpiration, 10*time.Minute)
}
