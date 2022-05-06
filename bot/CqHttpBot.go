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

type CqHttpBot struct{}

func (bot CqHttpBot) Config() register.BotConfig {
	return register.BotConfig{
		Name: "CqHttpBot",
	}
}

// Send 推送消息给CqHttpBot。
func (bot CqHttpBot) Send(crawlerResult [][]string, description string) error {
	var msg string

	for _, i := range crawlerResult {
		text := fmt.Sprintf("%s\n%s\n\n", i[1], i[0])
		msg += text
	}
	title := fmt.Sprintf("%s\n%s\n\n", description, utils.CurrentTime())

	client := utils.BotClient(Cfg.Bot.CqHttpBot.Timeout)

	allData := msg
	for { // 消息分片，避免过长消息发送失败
		if len(allData) > 3000 {
			left := allData[:3000]
			right := allData[3000:]
			index := strings.Index(right, "\n\n")
			left = left + right[:index]
			allData = right[index+2:]

			err := sendslip(client, left, title)
			if err != nil {
				return err
			}
		} else {
			err := sendslip(client, allData, title)
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func sendslip(client *http.Client, msg string, title string) error {
	msgEncode := url.QueryEscape(title + msg)
	data := fmt.Sprintf(`group_id=%d&message=%s`, Cfg.Bot.CqHttpBot.QQGroup, msgEncode)

	req, err := http.NewRequest("POST", Cfg.Bot.CqHttpBot.Api+"/send_group_msg?access_token="+Cfg.Bot.CqHttpBot.Key, strings.NewReader(data))
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
	fmt.Printf("[*] send to CqHttpBot: %s\n", respString)
	return nil
}
