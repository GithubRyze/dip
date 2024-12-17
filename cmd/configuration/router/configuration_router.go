package router

import "github.com/gin-gonic/gin"

func AddConfigurationRouter(v1 *gin.RouterGroup) {

	v1.DELETE("/configuration", deleteConfiguration)
	v1.POST("/configuration", createConfiguration)
}

func createConfiguration(ctx *gin.Context) {

}

func deleteConfiguration(ctx *gin.Context) {

}
