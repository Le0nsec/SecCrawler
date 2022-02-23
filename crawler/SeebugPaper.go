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

type SeebugPaper struct{}

func (crawler SeebugPaper) Config() register.CrawlerConfig {
	return register.CrawlerConfig{
		Name:        "SeebugPaper",
		Description: "SeebugPaper-安全技术精粹",
	}
}

// Get 获取Paper Seebug（知道创宇）前24小时内文章。
func (crawler SeebugPaper) Get() ([][]string, error) {
	client := utils.CrawlerClient()

	req, err := http.NewRequest("GET", "https://paper.seebug.org/rss/", nil)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(body)

	re := regexp.MustCompile(`<item><title>([\w\W]*?)</title><link>([\w\W]*?)</link><description>[\w\W]*?</description><pubDate>([\w\W]*?)</pubDate><guid>[\w\W]*?</guid><category>[\w\W]*?/category></item>`)
	result := re.FindAllStringSubmatch(strings.TrimSpace(bodyString), -1)

	var resultSlice [][]string
	fmt.Printf("[*] [SeebugPaper] crawler result:\n%s\n\n", utils.CurrentTime())
	for _, match := range result {
		utc, _ := time.LoadLocation("UTC")
		t, err := time.ParseInLocation(time.RFC1123Z, match[1:][2], utc)
		if err != nil {
			return nil, err
		}

		time_zone := time.FixedZone("CST", 8*3600)
		if !utils.IsIn24Hours(t.In(time_zone)) {
			// 默认时间顺序是从近到远
			break
		}

		// 去除title中的换行符
		re, _ = regexp.Compile(`\s{1,}`)
		match[1:][0] = re.ReplaceAllString(match[1:][0], "")

		fmt.Println(t.In(time_zone).Format("2006/01/02 15:04:05"))
		fmt.Println(match[1:][0])
		fmt.Printf("%s\n\n", match[1:][1])

		resultSlice = append(resultSlice, match[1:][0:2])
	}
	// slice中title和url调换位置，以符合统一的format
	for _, item := range resultSlice {
		item[0], item[1] = item[1], item[0]
	}
	if len(resultSlice) == 0 {
		return nil, errors.New("no records in the last 24 hours")
	}
	return resultSlice, nil
}
