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

// 垃圾桶路由
func TrashCanRouter(router *gin.Engine) {
	router.GET("api/trashcan/getDeviceInfo", api.GetDeviceInfo)                     // 获取垃圾桶设备全量信息		// 添加设备或修改信息时请求
	router.GET("api/trashcan/deviceStatus", api.DeviceStatus)                       // 获取设备状态信息			// 半小时
	router.GET("api/trashcan/deviceWarn", api.DeviceWarn)                           // 获取设备故障信息			// 半小时
	router.GET("api/trashcan/getConsumerScore", api.GetConsumerScore)               // 获取居民积分信息			// 每天一次
	router.GET("api/trashcan/getDeliveryRecord", api.GetDeliveryRecord)             // 获取前一天投递记录数据		// 每天一次
	router.GET("api/trashcan/getDeliveryRecordByDate", api.GetDeliveryRecordByDate) // 获取根据日期的投递记录数据
	router.GET("api/trashcan/getExchange", api.GetExchange)                         // 获取兑换记录数据			// 每天一次
	router.GET("api/trashcan/orgUnitList", api.OrgUnitList)                         // 获取小区数据				// 一星期一次
	router.GET("api/trashcan/getLogSweep", api.GetLogSweep)                         // 获取集中回收数据			// 每天一次
}
