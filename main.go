package main

import (
	_ "SecCrawler/bot"
	"SecCrawler/config"
	_ "SecCrawler/crawler"
	"SecCrawler/register"
	"SecCrawler/utils"
	"fmt"
	"log"
	"strings"

	"github.com/robfig/cron"
)

var cfg = config.GetGlobalConfig()

func main() {
	if !cfg.Debug {
		_cron := cron.New()
		spec := fmt.Sprintf("0 0 %d * * ?", cfg.CronTime)
		err := _cron.AddFunc(spec, start)
		// err := _cron.AddFunc("0 */1 * * * ?", start) //每分钟
		if err != nil {
			log.Fatalf("add cron error: %s\n", err.Error())
		}

		_cron.Start()
		defer _cron.Stop()
		select {}
	} else {
		start()
	}
}

func start() {
	fmt.Printf("%s\n[♥︎] crawler start at %s\n%s\n\n", strings.Repeat("-", 47), utils.CurrentTime(), strings.Repeat("-", 47))

	for crawlerName, crawler := range register.GetCrawlerMap() {
		crawlerResult, err := crawler.Get()
		if err != nil {
			log.Printf("crawl [%s] error: %s\n\n", crawlerName, err.Error())
			continue
		}
		for botName, bot := range register.GetBotMap() {
			err := bot.Send(crawlerResult, crawler.Config().Description)
			if err != nil {
				log.Printf("send [%s] to [%s] error: %s\n", crawlerName, botName, err.Error())
			}
		}
	}

}
