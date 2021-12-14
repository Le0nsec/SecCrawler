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

// sendHexQBot 推送消息给HexQBot。
func sendHexQBot(msg string) error {
	client := &http.Client{
		Timeout: time.Duration(cfg.HexQBot.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`{"msg": "%s", "num": %d}`, msg, cfg.HexQBot.QQGroup)
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
