package api

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var url = "http://www.iotcloud168.com:8111/intf/queryRealData"

// 获取气象实时数据的请求
func Atmosphere(c *gin.Context) {
	facId := c.Query("facId")
	a := getReal(facId)
	c.String(http.StatusOK, a)
}

func getReal(facId string) string {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url+"?facId="+facId, nil)
	if err != nil {
		log.Println("err")
	}
	// 添加请求头
	req.Header.Add("Content-type", "application/x-www-form-url")
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
	return string(b)
}

// 将字符串中的空格转为%20
func trans(input string) string {
	res := strings.Replace(input, " ", "%20", -1)
	return res
}
