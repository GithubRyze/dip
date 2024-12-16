package cmd

import (
	"dip/cmd/configuration"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddHttpRouter(c *gin.Engine) {
	checkHealthRouter(c)
	dipV1 := c.Group("/dip/v1/")
	configuration.AddConfigurationRouter(dipV1)
}

func checkHealthRouter(c *gin.Engine) {
	c.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
}
