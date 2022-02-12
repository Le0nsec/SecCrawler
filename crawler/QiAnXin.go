package crawler

import (
	"SecCrawler/register"
	"SecCrawler/utils"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type QiAnXin struct{}

func (crawler QiAnXin) Config() register.CrawlerConfig {
	return register.CrawlerConfig{
		Name:        "QiAnXin",
		Description: "奇安信攻防社区",
	}
}

// Get 获取奇安信前24小时内文章。
func (crawler QiAnXin) Get() ([][]string, error) {
	client := &http.Client{
		Timeout: time.Duration(4) * time.Second,
	}
	req, err := http.NewRequest("GET", "https://forum.butian.net/Rss", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(body)

	re := regexp.MustCompile(`<item><guid>([\w\W]*?)</guid><title>([\w\W]*?)</title><description>[\w\W]*?</description><source>[\w\W]*?</source><pubDate>([\w\W]*?)</pubDate></item>`)
	result := re.FindAllStringSubmatch(strings.TrimSpace(bodyString), -1)

	var resultSlice [][]string
	fmt.Printf("%s\n[*] [QiAnXin] crawler result:\n%s\n\n", strings.Repeat("-", 30), utils.CurrentTime())
	for _, match := range result {
		time_zone := time.FixedZone("CST", 8*3600)
		t, err := time.ParseInLocation("2006-01-02 15:04:05", match[1:][2], time_zone)
		if err != nil {
			return nil, err
		}

		if !utils.IsIn24Hours(t.In(time_zone)) {
			// 默认时间顺序是从近到远
			break
		}

		// 去除title中的换行符
		re, _ = regexp.Compile(`\s{1,}`)
		match[1:][0] = re.ReplaceAllString(match[1:][0], "")

		fmt.Println(t.In(time_zone).Format("2006/01/02 15:04:05"))
		fmt.Println(match[1:][1])
		fmt.Printf("%s\n\n", match[1:][0])

		resultSlice = append(resultSlice, match[1:][0:2])
	}
	if len(resultSlice) == 0 {
		return nil, errors.New("no records in the last 24 hours")
	}
	return resultSlice, nil

}
