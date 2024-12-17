package configuration

import (
	"dip/cmd/configuration/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func StartDipWebConfigurationServer(webPort string) error {
	ginEngine := gin.Default()
	addHttpRouter(ginEngine)
	svc := &http.Server{
		Addr:           ":" + webPort,
		Handler:        ginEngine,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return svc.ListenAndServe()
}

func addHttpRouter(c *gin.Engine) {
	checkHealthRouter(c)
	dipV1 := c.Group("/dip/v1/")
	router.AddConfigurationRouter(dipV1)
}

func checkHealthRouter(c *gin.Engine) {
	c.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
}
