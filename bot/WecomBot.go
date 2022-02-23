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

type WecomBot struct{}

func (bot WecomBot) Config() register.BotConfig {
	return register.BotConfig{
		Name: "WecomBot",
	}
}

// Send 推送消息给企业微信机器人。
func (bot WecomBot) Send(crawlerResult [][]string, description string) error {
	var msg string

	for _, i := range crawlerResult {
		text := fmt.Sprintf("> %s\\n\\n[%s](%s)\\n\\n\\n", i[1], i[0], i[0])
		msg += text
	}
	title := fmt.Sprintf("## %s\\n### %s\\n\\n\\n", description, utils.CurrentTime())

	client := utils.BotClient(Cfg.Bot.WecomBot.Timeout)

	data := fmt.Sprintf(`{"msgtype": "markdown", "markdown": {"content": "%s"}}`, title+msg)
	req, err := http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+Cfg.Bot.WecomBot.Key, strings.NewReader(data))
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
	fmt.Printf("[*] send to WecomBot: %s\n", respString)
	return nil
}
