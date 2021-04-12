package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (r Router) HandleNewWebsiteInteraction(c *gin.Context) {
	defer c.Done()

	websiteInteraction, err := NewWebsiteInteractionFromPost(c)
	if err != nil {
		log.Printf("error handling new website transaction: %s", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	exists, err := r.cacheSrv.CreateBoolEntryIfNotExists(websiteInteraction.UserPageUrlVisitKey())
	if err != nil {
		log.Printf("error creating new entry in cache: %s", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	if !exists {
		if err := r.cacheSrv.IncrementByKey(websiteInteraction.PageUrlKey(), 1); err != nil {
			log.Printf("error incrementing key in cache: %s", err.Error())
			c.Status(http.StatusBadRequest)
			return
		}
	}

	c.Status(http.StatusOK)
	return
}


func (r Router) HandleGetPageUniqueVisitorsCount(c *gin.Context) {
	defer c.Done()

	websiteInteraction, err := NewWebsiteInteractionFromGet(c.Request.URL.Query())
	if err != nil {
		log.Printf("error handling new get Page Unique Visitors Count request: %s", err.Error())
		c.String(http.StatusBadRequest, "")
		return
	}

	uniqueVisitors := r.cacheSrv.GetIntValueByKey(websiteInteraction.PageUrlKey())

	c.String(http.StatusOK, fmt.Sprintf("%d", uniqueVisitors))
	return
}