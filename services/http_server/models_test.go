package http_server

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestWebsiteInteraction(t *testing.T) {
	t1 := WebsiteInteraction{
		UserUid: "",
		PageUrl: "https://deus.ai",
		RequestedMethod: MethodPost,
	}
	err := t1.Validate()
	assert.Errorf(t, err, "user uid is a mandatory field")

	t1.RequestedMethod = MethodGet
	err2 := t1.Validate()
	assert.Nil(t, err2)

	t1.HandleUrl()
	urlB64:= base64.URLEncoding.EncodeToString([]byte("https://deus.ai"))
	assert.Equal(t, t1.PageUrlBase64,urlB64)
}
