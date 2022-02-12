package bot

import (
	"SecCrawler/register"
	"SecCrawler/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

	client := &http.Client{
		Timeout: time.Duration(cfg.HexQBot.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`{"msg": "%s", "num": %d, "key": "%s"}`, title+msg, cfg.HexQBot.QQGroup, cfg.HexQBot.Key)
	req, err := http.NewRequest("POST", cfg.HexQBot.Api, strings.NewReader(data))
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
