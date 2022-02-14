package bot

import (
	. "SecCrawler/config"
	"SecCrawler/register"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ServerChan struct{}

func (bot ServerChan) Config() register.BotConfig {
	return register.BotConfig{
		Name: "ServerChan",
	}
}

// Send 推送消息给Server酱。
func (bot ServerChan) Send(crawlerResult [][]string, description string) error {
	var msg string

	for _, i := range crawlerResult {
		text := fmt.Sprintf("%s\n[%s](%s)\n\n", i[1], i[0], i[0])
		msg += text
	}

	client := &http.Client{
		Timeout: time.Duration(Cfg.Bot.ServerChan.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`title=%s&desp=%s`, url.QueryEscape(description), url.QueryEscape(msg))

	req, err := http.NewRequest("POST", "https://sctapi.ftqq.com/"+Cfg.Bot.ServerChan.SendKey+".send", strings.NewReader(data))
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
	fmt.Printf("[*] send to ServerChan: %s\n", respString)
	return nil
}
