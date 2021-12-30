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
	//bytesData2 := fmt.Sprintf(`{"msgtype": "text","text": {"content":"%s"}`, msg)
	//fmt.Println(bytesData2)
	//client := &http.Client{}
	//data := make(map[string]interface{})
	//data["msgtype"] = "text"
	//data["text"] = "{\"content\":\"llllls\"}"
	//bytesData, _ := json.Marshal(data)
	//
	////req, _ := http.NewRequest("POST","http://httpbin.org/post",bytes.NewReader(bytesData))
	//req, _ := http.NewRequest("POST", "https://oapi.dingtalk.com/robot/send?access_token=629248e201b89c474b1fa1d543f359253faa1b673b3c977658126ceb087dbd55",bytes.NewReader(bytesData))
	//req.Header.Set("Content-type", "application/json")
	//
	//resp, _ := client.Do(req)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//return nil

	//data := fmt.Sprintf(`{"msgtype": "text","text": {"content":"%s"}}`, msg)
	client := &http.Client{
		Timeout: time.Duration(config.Cfg.HexQBot.Timeout) * time.Second,
	}

	data := fmt.Sprintf(`{"msgtype": "text","text": {"content":"%s"}}`, msg)
	fmt.Println("输入token值为：",config.Cfg.DingBot.Token)
	req, err := http.NewRequest("POST","https://oapi.dingtalk.com/robot/send?access_token=" + config.Cfg.DingBot.Token,strings.NewReader(data))
	//req, err := http.NewRequest("POST", "https://oapi.dingtalk.com/robot/send?access_token=629248e201b89c474b1fa1d543f359253faa1b673b3c977658126ceb087dbd55", strings.NewReader(data))
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
