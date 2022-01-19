package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// sendWecomBot 推送消息给企业微信机器人。
func sendWecomBot(msg string) error {
	client := &http.Client{
		Timeout: time.Duration(cfg.WecomBot.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`{"msgtype": "markdown", "markdown": {"content": "%s"}}`, msg)
	req, err := http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+cfg.WecomBot.Key, strings.NewReader(data))
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

// sendFeishuBot 推送消息给飞书群机器人。
func sendFeishuBot(msg string) error {
	client := &http.Client{
		Timeout: time.Duration(cfg.WecomBot.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`{"msg_type":"text","content":{"text":"%s"}}`, msg)
	req, err := http.NewRequest("POST", "https://open.feishu.cn/open-apis/bot/v2/hook/"+cfg.FeishuBot.Key, strings.NewReader(data))
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

// sendDingBot 推送消息给钉钉群机器人。
func sendDingBot(msg string) error {
	client := &http.Client{
		Timeout: time.Duration(cfg.DingBot.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`{"msgtype": "text","text": {"content":"%s"}}`, msg)
	req, err := http.NewRequest("POST", "https://oapi.dingtalk.com/robot/send?access_token="+cfg.DingBot.Token, strings.NewReader(data))
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

// sendHexQBot 推送消息给HexQBot。
func sendHexQBot(msg string) error {
	client := &http.Client{
		Timeout: time.Duration(cfg.HexQBot.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`{"msg": "%s", "num": %d, "key": "%s"}`, msg, cfg.HexQBot.QQGroup, cfg.HexQBot.Key)
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

// sendServerChan 推送消息给Server酱。
func sendServerChan(title, msg string) error {
	client := &http.Client{
		Timeout: time.Duration(cfg.ServerChan.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`title=%s&desp=%s`, url.QueryEscape(title), url.QueryEscape(msg))

	req, err := http.NewRequest("POST", "https://sctapi.ftqq.com/"+cfg.ServerChan.SendKey+".send", strings.NewReader(data))
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
