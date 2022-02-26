package bot

import (
	. "SecCrawler/config"
	"SecCrawler/register"
	"SecCrawler/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type WgpSecBot struct{}

func (bot WgpSecBot) Config() register.BotConfig {
	return register.BotConfig{
		Name: "WgpSecBot",
	}
}

// Send 推送消息给Server酱。
func (bot WgpSecBot) Send(crawlerResult [][]string, description string) error {
	var msg string

	for _, i := range crawlerResult {
		text := fmt.Sprintf("%s\n%s\n\n", i[1], i[0])
		msg += text
	}
	title := fmt.Sprintf("%s\n%s\n\n", description, utils.CurrentTime())

	client := utils.BotClient(Cfg.Bot.WgpSecBot.Timeout)

	data := fmt.Sprintf(`txt=%s`, url.QueryEscape(title+msg))

	req, err := http.NewRequest("POST", "https://api.bot.wgpsec.org/push/"+Cfg.Bot.WgpSecBot.Key, strings.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("[*] send to WgpSecBot: %s\n", respString)
	return nil
}
