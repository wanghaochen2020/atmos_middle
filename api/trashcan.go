package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	time "time"
)

var (
	// 所有路由的共同前缀，将内网服务的信息添加在后面
	trashCanUrl = "http://delanshi.net/dls/open/"
	// 设备厂商提供的key
	TrashCanKey = "55b2b2b9171fb22fde943b3a744f79f1"
	key         = "?orgId=" + TrashCanKey
)

// 获取垃圾桶设备全量信息
func GetDeviceInfo(c *gin.Context) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, trashCanUrl+"getDeviceInfo"+key, nil)
	if err != nil {
		log.Println("err")
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	c.String(http.StatusOK, string(b))
}

// 获取设备状态信息
func DeviceStatus(c *gin.Context) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, trashCanUrl+"deviceStatus"+key, nil)
	if err != nil {
		log.Println("err")
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	fmt.Printf("%s\n", string(b))
	c.String(http.StatusOK, string(b))
}

// 获取设备故障信息
func DeviceWarn(c *gin.Context) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, trashCanUrl+"deviceWarn"+key, nil)
	if err != nil {
		log.Println("err")
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	c.String(http.StatusOK, string(b))
}

// 获取居民积分信息
func GetConsumerScore(c *gin.Context) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, trashCanUrl+"getConsumerScore"+key, nil)
	if err != nil {
		log.Println("err")
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	c.String(http.StatusOK, string(b))
}

// 获取前一天投递记录数据
func GetDeliveryRecord(c *gin.Context) {
	u := trashCanUrl + "getDeliveryRecord" + key
	visitor := c.Query("type")
	if len(visitor) != 0 {
		// 如果传了参数则加在url后
		u = u + "&type=" + visitor
	} else {
		u = u + "&type=0"
	}
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		log.Println("err")
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	c.String(http.StatusOK, string(b))
}

// 获取根据日期的投递记录数据
func GetDeliveryRecordByDate(c *gin.Context) {
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	if len(startTime) == 0 {
		startTime = "2022-01-01 00:00:00"
	}
	if len(endTime) == 0 {
		endTime = time.Now().Format("2006-01-02 15:04:05")
	}
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, trashCanUrl+"getDeliveryRecordByDate"+key+"&startDate="+trans(startTime)+"&endDate="+trans(endTime), nil)
	if err != nil {
		log.Println("err")
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	c.String(http.StatusOK, string(b))
}

// 获取兑换记录数据
func GetExchange(c *gin.Context) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, trashCanUrl+"getExchange"+key, nil)
	if err != nil {
		log.Println("err")
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	c.String(http.StatusOK, string(b))
}

// 获取小区数据
func OrgUnitList(c *gin.Context) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, trashCanUrl+"orgUnitList"+key, nil)
	if err != nil {
		log.Println("err")
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	c.String(http.StatusOK, string(b))
}

// 获取集中回收数据
func GetLogSweep(c *gin.Context) {
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	if len(startTime) == 0 {
		startTime = "2000-01-01 00:00:00"
	}
	if len(endTime) == 0 {
		endTime = time.Now().Format("2006-01-02 15:04:05")
	}
	u := trashCanUrl + "getLogSweep" + key + "&startTime=" + trans(startTime) + "&endTime=" + trans(endTime)
	kind := c.Query("type")
	if len(kind) != 0 {
		u = u + "&type=" + kind
	} else {
		u = u + "&type=-1"
	}
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		log.Println("err")
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	c.String(http.StatusOK, string(b))
}
