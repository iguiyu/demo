package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusNoContent(c *gin.Context) {
	c.Data(http.StatusNoContent, gin.MIMEJSON, nil)
}

func SetQuery(c *gin.Context, key, value string) {
	c.Request.URL.RawQuery += fmt.Sprintf("&%s=%s", key, value)
}
