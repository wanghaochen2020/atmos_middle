package routers

import (
	"github.com/gin-gonic/gin"
	"middle/api"
)

// 气象站路由
func AtmosphereDeviceRouter(router *gin.Engine) {
	router.GET("api/atmosphere", api.Atmosphere)
	router.GET("api/history", api.AtmosphereHistory)
}
