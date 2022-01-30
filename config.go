package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Debug        bool   `yaml:"Debug" binding:"required"`
	CronTime     uint8  `yaml:"CronTime" binding:"required"`
	ChromeDriver string `yaml:"ChromeDriver"`

	WecomBot   *WecomBotStruct   `yaml:"WecomBot"`
	FeishuBot  *FeishuBotStruct  `yaml:"FeishuBot"`
	DingBot    *DingBotStruct    `yaml:"DingBot"`
	HexQBot    *HexQBotStruct    `yaml:"HexQBot"`
	ServerChan *ServerChanStruct `yaml:"ServerChan"`

	EdgeForum   *EdgeForumStruct   `yaml:"EdgeForum"`
	XianZhi     *XianZhiStruct     `yaml:"XianZhi"`
	SeebugPaper *SeebugPaperStruct `yaml:"SeebugPaper"`
	Anquanke    *AnquankeStruct    `yaml:"Anquanke"`
	Tttang      *TttangStruct      `yaml:"Tttang"`
	QiAnXin     *QiAnXinStruct     `yaml:"QiAnXin"`
}

type WecomBotStruct struct {
	Enabled bool   `yaml:"enabled"`
	Key     string `yaml:"key"`
	Timeout uint8  `yaml:"timeout"`
}

type FeishuBotStruct struct {
	Enabled bool   `yaml:"enabled"`
	Key     string `yaml:"key"`
	Timeout uint8  `yaml:"timeout"`
}

type DingBotStruct struct {
	Enabled bool   `yaml:"enabled"`
	Token   string `yaml:"token"`
	Timeout uint8  `yaml:"timeout"`
}

type HexQBotStruct struct {
	Enabled bool   `yaml:"enabled"`
	Api     string `yaml:"api"`
	QQGroup uint64 `yaml:"qqgroup"`
	Key     string `yaml:"key"`
	Timeout uint8  `yaml:"timeout"`
}

type ServerChanStruct struct {
	Enabled bool   `yaml:"enabled"`
	SendKey string `yaml:"sendkey"`
	Timeout uint8  `yaml:"timeout"`
}

type EdgeForumStruct struct {
	Enabled bool `yaml:"enabled"`
}

type XianZhiStruct struct {
	Enabled bool `yaml:"enabled"`
}

type SeebugPaperStruct struct {
	Enabled bool `yaml:"enabled"`
}

type AnquankeStruct struct {
	Enabled bool `yaml:"enabled"`
}

type TttangStruct struct {
	Enabled bool `yaml:"enabled"`
}

type QiAnXinStruct struct {
	Enabled bool `yaml:"enablesd"`
}

// 全局Config
var cfg *Config

var siteDescriptionMap = map[string]string{
	"EdgeForum":   "棱角社区攻防日报",
	"XianZhi":     "先知安全技术社区",
	"SeebugPaper": "SeebugPaper-安全技术精粹",
	"Anquanke":    "安全客-安全资讯平台",
	"Tttang":      "跳跳糖-安全与分享社区",
	"QiAnXin":     "奇安信攻防社区",
}

func init() {
	log.SetPrefix("[!] ")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yml")
	// 判断config文件是否存在，不存在则初始化
	if _, err := os.Stat("./config.yml"); os.IsNotExist(err) {
		fmt.Println("[*] config.yml does not exist!")
		f, err := os.Create("./config.yml")
		if err != nil {
			log.Fatalf("create config file error: %s\n", err.Error())
		}
		defer f.Close()
		configString := `
############### CronSetting ###############
# 开启则一次性爬取后退出程序
Debug: false
# 设置每天整点爬取推送时间，范围 0 ~ 23（整数）
CronTime: 11
# 设置Selenium使用的ChromeDriver路径，支持相对路径或绝对路径（如果不爬取先知社区可以不用设置）
ChromeDriver: ./chromedriver/linux64


############### BotConfig ###############

# 企业微信群机器人
# https://work.weixin.qq.com/api/doc/90000/90136/91770
WecomBot:
  enabled: false
  key: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
  timeout: 2  # second

# 飞书群机器人
# https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
FeishuBot:
  enabled: false
  key: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
  timeout: 2

# 钉钉群机器人
# https://open.dingtalk.com/document/robots/custom-robot-access
DingBot:
  enabled: false
  token: xxxxxxxxxxxxxxxxxxxxxx
  timeout: 2

# HexQBot
# https://github.com/Am473ur/HexQBot
HexQBot:
  enabled: false
  api: http://xxxxxx.com/send
  qqgroup: 000000000
  key: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
  timeout: 2

# Server酱
# https://sct.ftqq.com/
ServerChan:
  enabled: false
  sendkey: xxxxxxxxxxxxxxxxxxxxxx
  timeout: 2


############### SiteEnable ###############

# 棱角社区
# https://forum.ywhack.com/forum-59-1.html
EdgeForum:
  enabled: true

# 先知安全技术社区
# https://xz.aliyun.com/
XianZhi:
  enabled: true

# SeebugPaper（知道创宇404实验室）
# https://paper.seebug.org/
SeebugPaper:
  enabled: true

# 安全客
# https://www.anquanke.com/
Anquanke:
  enabled: true

# 跳跳糖
# http://tttang.com/
Tttang:
  enable: true

# 奇安信攻防社区
# https://forum.butian.net/community/all/newest
QiAnXin:
  enabled: true
`
		_, err = f.WriteString(configString)
		if err != nil {
			log.Fatalf("write config file error: %s\n", err.Error())
		}
		f.Sync()
		fmt.Println("[*] The configuration file has been initialized.")
		os.Exit(0)
	} else {
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("read config file error: %s\n", err.Error())
		}

		err = viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatalf("unmarshal config error: %s\n", err.Error())
		}
		fmt.Printf("[*] load config success!\n\n")
	}
}
