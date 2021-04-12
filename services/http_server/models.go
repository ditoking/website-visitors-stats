package http_server

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
	"net/url"
)

const (
	MethodGeneric = iota
	MethodPost
	MethodGet
)

type WebsiteInteraction struct {
	PageUrl         string `json:"page_url" form:"page_url" schema:"page_url"`
	PageUrlBase64   string `json:"page_url_base_64"`
	UserUid         string `json:"user_uid" form:"user_uid"`
	RequestedMethod int
}

func NewWebsiteInteractionFromPost(c *gin.Context) (*WebsiteInteraction, error) {
	websiteInteraction := new(WebsiteInteraction)

	if err := c.ShouldBind(&websiteInteraction); err != nil {
		return nil, err
	}

	websiteInteraction.RequestedMethod = MethodPost

	if err := websiteInteraction.Validate(); err != nil {
		return nil, err
	}

	websiteInteraction.HandleUrl()

	return websiteInteraction, nil
}

func NewWebsiteInteractionFromGet(queryParams map[string][]string) (*WebsiteInteraction, error) {
	websiteInteraction := new(WebsiteInteraction)
	decoder := schema.NewDecoder()

	if err := decoder.Decode(websiteInteraction, queryParams); err != nil {
		return nil, err
	}

	websiteInteraction.RequestedMethod = MethodGet

	if err := websiteInteraction.Validate(); err != nil {
		return nil, err
	}

	websiteInteraction.HandleUrl()

	return websiteInteraction, nil
}

func (wi *WebsiteInteraction) Validate() error {
	if wi.UserUid == "" && wi.RequestedMethod == MethodPost {
		return errors.New("user uid is a mandatory field")
	}

	_, err := url.Parse(wi.PageUrl)
	if err != nil {
		return err
	}

	return nil
}

func (wi *WebsiteInteraction) HandleUrl() {
	wi.PageUrlBase64 = base64.URLEncoding.EncodeToString([]byte(wi.PageUrl))
}

func (wi *WebsiteInteraction) UserPageUrlVisitKey() string {
	return fmt.Sprintf("unique_visit:%s_%s", wi.PageUrlBase64, wi.UserUid)
}

func (wi *WebsiteInteraction) PageUrlKey() string {
	return fmt.Sprintf("page:%s", wi.PageUrlBase64)
}
