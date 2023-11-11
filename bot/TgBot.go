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

type TgBot struct{}

func (bot TgBot) Config() register.BotConfig {
	return register.BotConfig{
		Name: "TgBot",
	}
}

// Send 推送消息给telegram api。
func (bot TgBot) Send(crawlerResult [][]string, description string) error {
	var msg string

	for _, i := range crawlerResult {
		text := fmt.Sprintf("%s\n[%s](%s)\n\n", i[1], i[0], i[0])
		msg += text
	}
	title := fmt.Sprintf("%s\n%s\n\n", description, utils.CurrentTime())

	client := utils.BotClient(Cfg.Bot.TgBot.Timeout)

	data := fmt.Sprintf("{\"chat_id\":\"%s\",\"text\":\"%s\",\"parse_mode\":\"MarkdownV2\"}", Cfg.Bot.TgBot.ChatId, title+msg)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.telegram.org/%s/sendMessage", Cfg.Bot.TgBot.Token), strings.NewReader(data))
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
	fmt.Printf("[*] send to TgBot: %s\n", respString)
	return nil
}
