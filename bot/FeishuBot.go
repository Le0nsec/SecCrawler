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

type FeishuBot struct{}

func (bot FeishuBot) Config() register.BotConfig {
	return register.BotConfig{
		Name: "FeishuBot",
	}
}

// Send 推送消息给飞书群机器人。
func (bot FeishuBot) Send(crawlerResult [][]string, description string) error {
	var msg string

	for _, i := range crawlerResult {
		text := fmt.Sprintf("%s\\n%s\\n\\n", i[1], i[0])
		msg += text
	}
	title := fmt.Sprintf("%s\\n%s\\n\\n", description, utils.CurrentTime())

	client := utils.BotClient(Cfg.Bot.FeishuBot.Timeout)

	data := fmt.Sprintf(`{"msg_type":"text","content":{"text":"%s"}}`, title+msg)
	req, err := http.NewRequest("POST", "https://open.feishu.cn/open-apis/bot/v2/hook/"+Cfg.Bot.FeishuBot.Key, strings.NewReader(data))
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
	fmt.Printf("[*] send to FeishuBot: %s\n", respString)
	return nil
}
