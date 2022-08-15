package receive

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"middle/dataReceive/def"
	"net/http"
	"strconv"
)

const (
	baseUrl                    = "http://43.138.78.252:7777/"
	getDeviceInfoUrl           = baseUrl + "api/trashcan/getDeviceInfo"
	deviceStatusUrl            = baseUrl + "api/trashcan/deviceStatus"
	deviceWarnUrl              = baseUrl + "api/trashcan/deviceWarn"
	getConsumerScoreUrl        = baseUrl + "api/trashcan/getConsumerScore"
	getDeliveryRecordUrl       = baseUrl + "api/trashcan/getDeliveryRecord"
	getDeliveryRecordByDateUrl = baseUrl + "api/trashcan/getDeliveryRecordByDate"
	getExchangeUrl             = baseUrl + "api/trashcan/getExchange"
	orgUnitListUrl             = baseUrl + "api/trashcan/orgUnitList"
	getLogSweepUrl             = baseUrl + "api/trashcan/getLogSweep"
)

// 根据url从中转服务器获取设备数据字符串
func getData(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	n, _ := ioutil.ReadAll(resp.Body)
	return string(n)
}

// 获取设备信息全量数据
func GetDeviceInfo() []def.TrashCanInfo {
	data := getData(getDeviceInfoUrl)
	var res []def.TrashCanInfo
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Println("unmarshal fail")
	}
	return res
}

// 获取设备状态数据
func GetDeviceStatus() []def.TrashCanState {
	data := getData(deviceStatusUrl)
	var res []def.TrashCanState
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Println("unmarshal fail")
	}
	return res
}

// 获取设备故障数据
func GetDeviceWarn() []def.TrashCanFault {
	data := getData(deviceWarnUrl)
	var res []def.TrashCanFault
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Println("unmarshal fail")
	}
	return res
}

// 获取居民积分数据
func GetConsumerScore() []def.TrashCanScore {
	data := getData(getConsumerScoreUrl)
	var res []def.TrashCanScore
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Println("unmarshal fail")
	}
	return res
}

// 获取前一天投递记录数据
func GetDeliveryRecord(queryType int) []def.TrashCanDelivery {
	query := strconv.Itoa(queryType)
	data := getData(getDeliveryRecordUrl + "?type=" + query)
	var res []def.TrashCanDelivery
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Println("unmarshal fail")
	}
	return res
}

// 根据日期获取投递记录数据
func GetDeliveryRecordByDate(startTime, endTime string) []def.TrashCanDelivery {
	data := getData(getDeliveryRecordByDateUrl + "?startTime=" + startTime + "&endTime=" + endTime)
	var res []def.TrashCanDelivery
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Println("unmarshal fail")
	}
	return res
}

// 获取兑换记录数据
func GetExchange() []def.TrashCanExchange {
	data := getData(getExchangeUrl)
	var res []def.TrashCanExchange
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Println("unmarshal fail")
	}
	return res
}

// 获取小区数据
func GetUnitList() def.UnitData {
	data := getData(orgUnitListUrl)
	var res def.UnitData
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Println("unmarshal fail")
	}
	return res
}

// 获取集中回收数据
func GetLogSweep(startTime, endTime string, queryType int) []def.Recycle {
	query := strconv.Itoa(queryType)
	data := getData(getLogSweepUrl + "?startTime=" + startTime + "&endTime=" + endTime + "&type=" + query)
	var res []def.Recycle
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Println("unmarshal fail")
	}
	return res
}
