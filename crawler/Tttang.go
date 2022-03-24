package crawler

import (
	"SecCrawler/register"
	"SecCrawler/utils"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

type Tttang struct{}

func (crawler Tttang) Config() register.CrawlerConfig {
	return register.CrawlerConfig{
		Name:        "Tttang",
		Description: "跳跳糖-安全与分享社区",
	}
}

// Get 获取跳跳糖前24小时内文章。
func (crawler Tttang) Get() ([][]string, error) {
	client := utils.CrawlerClient()

	req, err := http.NewRequest("GET", "http://tttang.com/rss.xml", nil)
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

	fp := gofeed.NewParser()
	feed, err := fp.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var resultSlice [][]string
	fmt.Printf("[*] [Tttang] crawler result:\n%s\n\n", utils.CurrentTime())

	for _, item := range feed.Items {
		time_zone := time.FixedZone("CST", 8*3600)
		t, err := time.ParseInLocation(time.RFC1123Z, item.Published, time_zone)
		if err != nil {
			return nil, err
		}

		if !utils.IsIn24Hours(t.In(time_zone)) {
			// 默认时间顺序是从近到远
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
