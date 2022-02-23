package bot

import (
	. "SecCrawler/config"
	"SecCrawler/register"
	"SecCrawler/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type HexQBot struct{}

func (bot HexQBot) Config() register.BotConfig {
	return register.BotConfig{
		Name: "HexQBot",
	}
}

// Send 推送消息给HexQBot。
func (bot HexQBot) Send(crawlerResult [][]string, description string) error {
	var msg string

	for _, i := range crawlerResult {
		text := fmt.Sprintf("%s\\n%s\\n\\n", i[1], i[0])
		msg += text
	}
	title := fmt.Sprintf("%s\\n%s\\n\\n", description, utils.CurrentTime())

	client := utils.BotClient(Cfg.Bot.HexQBot.Timeout)

	data := fmt.Sprintf(`{"msg": "%s", "num": %d, "key": "%s"}`, title+msg, Cfg.Bot.HexQBot.QQGroup, Cfg.Bot.HexQBot.Key)
	req, err := http.NewRequest("POST", Cfg.Bot.HexQBot.Api, strings.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("[*] send to HexQBot: %s\n", respString)
	return nil
}
