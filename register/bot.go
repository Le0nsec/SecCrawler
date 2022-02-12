package register

import "fmt"

type BotConfig struct {
	Name string // Bot名称
}

type Bot interface {
	Config() BotConfig                                       // Bot名称
	Send(crawlerResult [][]string, description string) error // 推送方法
}

var botMap = map[string]Bot{}

func RegisterBot(bot Bot) {
	fmt.Printf("[+] register bot: [%s]\n", bot.Config().Name)
	botMap[bot.Config().Name] = bot
}

func GetBotMap() map[string]Bot {
	return botMap
}
