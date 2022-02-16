package utils

import (
	"SecCrawler/config"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SITE_NOT_FOUND    = 4000
	ARTICLE_NOT_FOUND = 4001
	INVALID_AUTH_KEY  = 4002
)

func CurrentTime() string {
	time_zone := time.FixedZone("CST", 8*3600) // 8*3600 = 8h
	n := time.Now().In(time_zone)
	// 获取时间，格式如2006/01/02 15:04:05
	t := n.Format("2006/01/02 15:04:05")
	weekMap := map[time.Weekday]string{0: "星期日", 1: "星期一", 2: "星期二", 3: "星期三", 4: "星期四", 5: "星期五", 6: "星期六"}
	formatTime := fmt.Sprintf("%s %s", t, weekMap[n.Weekday()])
	return formatTime
}

func IsIn24Hours(t time.Time) bool {
	time_zone := time.FixedZone("CST", 8*3600) // 8*3600 = 8h
	now := time.Now().In(time_zone)
	// 根据config生成每日整点时间
	var hour int
	if config.Test {
		hour = now.Hour()
	} else {
		hour = int(config.Cfg.Cron.Time)
	}
	cronTime := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, time_zone)
	subTime := cronTime.Sub(t)
	if subTime > time.Duration(24)*time.Hour || subTime < time.Duration(0) {
		return false
	}
	return true
}

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ErrorResp(c *gin.Context, code int, err error) {
	log.Printf("Resp error: [%s]", err.Error())
	c.JSON(200, Resp{
		Code: code,
		Msg:  err.Error(),
		Data: nil,
	})
	c.Abort()
}

func ErrorStrResp(c *gin.Context, code int, str string) {
	log.Printf("Resp error: [%s]", str)
	c.JSON(200, Resp{
		Code: code,
		Msg:  str,
		Data: nil,
	})
	c.Abort()
}

func SuccessResp(c *gin.Context, data ...interface{}) {
	if len(data) == 0 {
		c.JSON(200, Resp{
			Code: 200,
			Msg:  "success",
			Data: nil,
		})
		return
	}
	c.JSON(200, Resp{
		Code: 200,
		Msg:  "success",
		Data: data[0],
	})
}
