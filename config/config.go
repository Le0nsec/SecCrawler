package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

const Banner = `
  _____            _____                    _           
 / ____|          / ____|                  | |          
| (___   ___  ___| |     _ __ __ ___      _| | ___ _ __ 
 \___ \ / _ \/ __| |    | '__/ _  \ \ /\ / / |/ _ \ '__|
 ____) |  __/ (__| |____| | | (_| |\ V  V /| |  __/ |   
|_____/ \___|\___|\_____|_|  \__,_| \_/\_/ |_|\___|_|   																									  
`

// 全局Config
var Cfg *Config

var (
	Test       bool
	Version    bool
	Help       bool
	Generate   bool
	ConfigFile string

	GITHUB    string = "https://github.com/Le0nsec/SecCrawler"
	TAG       string = "v2.4"
	GOVERSION string = "go1.21.3"
)

func DefaultConfig() Config {
	return Config{
		ChromeDriver: "./chromedriver/linux64",
		Proxy: ProxyStruct{
			ProxyUrl:            "http://127.0.0.1:7890",
			CrawlerProxyEnabled: false,
			BotProxyEnabled:     false,
		},
		Cron: CronStruct{
			Enabled: false,
			Time:    11,
		},
		Api: ApiStruct{
			Enabled: false,
			Debug:   false,
			Host:    "127.0.0.1",
			Port:    8080,
			Auth:    "auth_key_here",
		},
		Crawler: CrawlerStruct{
			EdgeForum:   EdgeForumStruct{Enabled: false},
			XianZhi:     XianZhiStruct{Enabled: false, UseChromeDriver: true, CustomRSSURL: ""},
			SeebugPaper: SeebugPaperStruct{Enabled: false},
			Anquanke:    AnquankeStruct{Enabled: false},
			Tttang:      TttangStruct{Enabled: false},
			QiAnXin:     QiAnXinStruct{Enabled: false},
			// DongJian:    DongJianStruct{Enabled: false},
			Lab: LabStruct{
				Enabled:     false,
				NoahLab:     NoahLabStruct{Enabled: true},
				Blog360:     Blog360Struct{Enabled: true},
				Nsfocus:     NsfocusStruct{Enabled: true},
				Xlab:        XlabStruct{Enabled: true},
				AlphaLab:    AlphaLabStruct{Enabled: true},
				Netlab:      NetlabStruct{Enabled: true},
				RiskivyBlog: RiskivyBlogStruct{Enabled: true},
				TSRCBlog:    TSRCBlogStruct{Enabled: true},
				X1cT34m:     X1cT34mStruct{Enabled: true},
			},
			HuoxianZone: HuoxianZoneStruct{
				Enabled: true,
			},
		},
		Bot: BotStruct{
			WecomBot: WecomBotStruct{
				Enabled: false,
				Key:     "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				Timeout: 2,
			},
			FeishuBot: FeishuBotStruct{
				Enabled: false,
				Key:     "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				Timeout: 2,
			},
			DingBot: DingBotStruct{
				Enabled: false,
				Token:   "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
				Timeout: 2,
			},
			HexQBot: HexQBotStruct{
				Enabled: false,
				Api:     "http://xxxxxx.com/send",
				QQGroup: 000000000,
				Key:     "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				Timeout: 2,
			},
			ServerChan: ServerChanStruct{
				Enabled: false,
				SendKey: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
				Timeout: 2,
			},
			WgpSecBot: WgpSecBotStruct{
				Enabled: false,
				Key:     "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
				Timeout: 2,
			},
			TgBot: TgBotStruct{
				Enabled: false,
				Token:   "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
				ChatId:  "@channelusername",
				Timeout: 2,
			},
		},
	}
}

func configToYaml() string {
	b, err := yaml.Marshal(DefaultConfig())
	if err != nil {
		log.Fatalf("unable to marshal config to yaml: %s", err.Error())
	}
	return string(b)
}

func ConfigInit() {
	log.SetPrefix("[!] ")
	viper.SetConfigType("yaml")
	viper.SetConfigFile(ConfigFile)
	// 判断config文件是否存在
	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		if Generate {
			f, err := os.Create(ConfigFile)
			if err != nil {
				log.Fatalf("create config file error: %s\n", err.Error())
			}
			defer f.Close()

			_, err = f.WriteString(configToYaml())
			if err != nil {
				log.Fatalf("write config file error: %s\n", err.Error())
			}
			f.Sync()
			fmt.Println("[*] The configuration file has been initialized.")
			os.Exit(0)
		} else {
			fmt.Println("[!] The configuration file does not exist, please use `-init`")
			os.Exit(0)
		}
	} else {
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("read config file error: %s\n", err.Error())
		}

		err = viper.Unmarshal(&Cfg)
		if err != nil {
			log.Fatalf("unmarshal config error: %s\n", err.Error())
		}
		fmt.Printf("[*] load config success!\n\n")
	}
}
