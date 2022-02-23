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

type DingBot struct{}

func (bot DingBot) Config() register.BotConfig {
	return register.BotConfig{
		Name: "DingBot",
	}
}

// Send 推送消息给钉钉群机器人。
func (bot DingBot) Send(crawlerResult [][]string, description string) error {
	var msg string

	for _, i := range crawlerResult {
		text := fmt.Sprintf("%s\\n%s\\n\\n", i[1], i[0])
		msg += text
	}
	title := fmt.Sprintf("%s\\n%s\\n\\n", description, utils.CurrentTime())

	client := utils.BotClient(Cfg.Bot.DingBot.Timeout)

	data := fmt.Sprintf(`{"msgtype": "text","text": {"content":"%s"}}`, title+msg)
	req, err := http.NewRequest("POST", "https://oapi.dingtalk.com/robot/send?access_token="+Cfg.Bot.DingBot.Token, strings.NewReader(data))
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
	fmt.Printf("[*] send to DingBot: %s\n", respString)
	return nil
}
