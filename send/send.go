package send

import (
	"SecCrawler/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	//"honnef.co/go/tools/analysis/facts/nilness"
)

// SendWecomBot 推送消息给企业微信机器人。
func SendWecomBot(msg string) error {
	client := &http.Client{
		Timeout: time.Duration(config.Cfg.WecomBot.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`{"msgtype": "markdown", "markdown": {"content": "%s"}}`, msg)
	req, err := http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+config.Cfg.WecomBot.Key, strings.NewReader(data))
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

// SendFeishuBot 推送消息给飞书群机器人。
func SendFeishuBot(msg string) error {
	client := &http.Client{
		Timeout: time.Duration(config.Cfg.WecomBot.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`{"msg_type":"text","content":{"text":"%s"}}`, msg)
	req, err := http.NewRequest("POST", "https://open.feishu.cn/open-apis/bot/v2/hook/"+config.Cfg.FeishuBot.Key, strings.NewReader(data))
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

// SendDingBot 推送消息给钉钉群机器人。
func SendDingBot(msg string) error {
	client := &http.Client{
		Timeout: time.Duration(config.Cfg.DingBot.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`{"msgtype": "text","text": {"content":"%s"}}`, msg)
	req, err := http.NewRequest("POST", "https://oapi.dingtalk.com/robot/send?access_token="+config.Cfg.DingBot.Token, strings.NewReader(data))
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

// SendHexQBot 推送消息给HexQBot。
func SendHexQBot(msg string) error {
	client := &http.Client{
		Timeout: time.Duration(config.Cfg.HexQBot.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`{"msg": "%s", "num": %d}`, msg, config.Cfg.HexQBot.QQGroup)
	req, err := http.NewRequest("POST", config.Cfg.HexQBot.Api, strings.NewReader(data))
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

// SendServerChan 推送消息给Server酱。
func SendServerChan(title, msg string) error {
	client := &http.Client{
		Timeout: time.Duration(config.Cfg.ServerChan.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`title=%s&desp=%s`, url.QueryEscape(title), url.QueryEscape(msg))

	req, err := http.NewRequest("POST", "https://sctapi.ftqq.com/"+config.Cfg.ServerChan.SendKey+".send", strings.NewReader(data))
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
