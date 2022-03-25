package lab

import (
	"SecCrawler/register"
	"SecCrawler/utils"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

type RiskivyBlog struct{}

func (crawler RiskivyBlog) Config() register.CrawlerConfig {
	return register.CrawlerConfig{
		Name:        "Lab.RiskivyBlog",
		Description: "斗象能力中心",
	}
}

// Get 获取 RiskivyBlog 前24小时内文章。
func (crawler RiskivyBlog) Get() ([][]string, error) {
	client := utils.CrawlerClient()

	req, err := http.NewRequest("GET", "https://blog.riskivy.com/feed/", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fp := gofeed.NewParser()
	feed, err := fp.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var resultSlice [][]string
	fmt.Printf("[*] [RiskivyBlog] crawler result:\n%s\n\n", utils.CurrentTime())

	for _, item := range feed.Items {
		t, err := time.Parse(time.RFC1123Z, item.Published)
		if err != nil {
			return nil, err
		}

		time_zone := time.FixedZone("CST", 8*3600)
		if !utils.IsIn24Hours(t.In(time_zone)) {
			break
		}

		fmt.Println(t.In(time_zone).Format("2006/01/02 15:04:05"))
		fmt.Println(item.Title)
		fmt.Printf("%s\n\n", item.Link)

		var s []string
		s = append(s, item.Link, item.Title)
		resultSlice = append(resultSlice, s)
	}

	if len(resultSlice) == 0 {
		return nil, errors.New("no records in the last 24 hours")
	}
	return resultSlice, nil
}
