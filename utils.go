package main

import (
	"fmt"
	"time"
)

func currentTime() string {
	time_zone := time.FixedZone("CST", 8*3600) // 8*3600 = 8h
	n := time.Now().In(time_zone)
	// 获取时间，格式如2006/01/02 15:04:05
	t := n.Format("2006/01/02 15:04:05")
	weekMap := map[time.Weekday]string{0: "星期日", 1: "星期一", 2: "星期二", 3: "星期三", 4: "星期四", 5: "星期五", 6: "星期六"}
	formatTime := fmt.Sprintf("%s %s", t, weekMap[n.Weekday()])
	return formatTime
}

func isIn24Hours(t time.Time) bool {
	time_zone := time.FixedZone("CST", 8*3600) // 8*3600 = 8h
	now := time.Now().In(time_zone)
	// 根据config生成每日整点时间
	cronTime := time.Date(now.Year(), now.Month(), now.Day(), int(cfg.CronTime), 0, 0, 0, time_zone)
	subTime := cronTime.Sub(t)
	if subTime > time.Duration(24)*time.Hour || subTime < time.Duration(0) {
		return false
	}
	return true
}

// wecomBotFormat 格式化消息为markdown格式。
func wecomBotFormat(crawlerResult [][]string, site string) (msg string) {
	for _, i := range crawlerResult {
		text := fmt.Sprintf("> %s\\n\\n[%s](%s)\\n\\n\\n", i[1], i[0], i[0])
		msg += text
	}
	title := fmt.Sprintf("## %s\\n### %s\\n\\n\\n", siteDescriptionMap[site], currentTime())
	return title + msg
}

// FeishuBotFormat 格式化消息。
func FeishuBotFormat(crawlerResult [][]string, site string) (msg string) {
	for _, i := range crawlerResult {
		text := fmt.Sprintf("%s\\n%s\\n\\n", i[1], i[0])
		msg += text
	}
	title := fmt.Sprintf("%s\\n%s\\n\\n", siteDescriptionMap[site], currentTime())
	return title + msg
}

// hexQBotFormat 格式化消息。
func hexQBotFormat(crawlerResult [][]string, site string) (msg string) {
	for _, i := range crawlerResult {
		text := fmt.Sprintf("%s\\n%s\\n\\n", i[1], i[0])
		msg += text
	}
	title := fmt.Sprintf("%s\\n%s\\n\\n", siteDescriptionMap[site], currentTime())
	return title + msg
}

// serverChanFormat 格式化消息。
func serverChanFormat(crawlerResult [][]string, site string) (title, msg string) {
	for _, i := range crawlerResult {
		text := fmt.Sprintf("%s\n[%s](%s)\n\n", i[1], i[0], i[0])
		msg += text
	}
	return siteDescriptionMap[site], msg
}
