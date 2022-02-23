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

type Anquanke struct{}

func (crawler Anquanke) Config() register.CrawlerConfig {
	return register.CrawlerConfig{
		Name:        "Anquanke",
		Description: "安全客-安全资讯平台",
	}
}

// Get 获取安全客前24小时内文章。
func (crawler Anquanke) Get() ([][]string, error) {
	client := utils.CrawlerClient()

	req, err := http.NewRequest("GET", "https://www.anquanke.com/knowledge", nil)
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

	//去除STYLE
	re, _ := regexp.Compile(`\<style[\S\s]+?\</style\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除SCRIPT
	re, _ = regexp.Compile(`\<script[\S\s]+?\</script\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除head
	re, _ = regexp.Compile(`\<head[\S\s]+?\</head\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除header
	re, _ = regexp.Compile(`\<header[\S\s]+?\</header\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除footer
	re, _ = regexp.Compile(`\<footer[\S\s]+?\</footer\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除sidebar
	re, _ = regexp.Compile(`\<div class="load-more[\S\s]+?\</html\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除tag及desp
	re, _ = regexp.Compile(`\<div class="tags  hide-in-mobile-device[\S\s]+?class="fa fa-clock-o"\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除多余的信息
	re, _ = regexp.Compile(`\<div class="article-item common-item">[\S\s]+?common-item-right"\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除连续的换行符
	re, _ = regexp.Compile(`\s{2,}`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除href中可能存在的class
	re, _ = regexp.Compile(`class="red-title"`)
	bodyString = re.ReplaceAllString(bodyString, "")

	re = regexp.MustCompile(`<div class="title"><a target="_blank" rel="noopener noreferrer"href="(.*?)"> (.*?)</a></div></i>(.*?)</span>`)
	result := re.FindAllStringSubmatch(strings.TrimSpace(bodyString), -1)

	var resultSlice [][]string
	fmt.Printf("[*] [Anquanke] crawler result:\n%s\n\n", utils.CurrentTime())
	for _, match := range result {
		match[1:][0] = "https://www.anquanke.com" + match[1:][0]
		time_zone := time.FixedZone("CST", 8*3600)
		t, err := time.ParseInLocation("2006-01-02 15:04:05", match[1:][2], time_zone)
		if err != nil {
			return nil, err
		}

		if !utils.IsIn24Hours(t) {
			// 默认时间顺序是从近到远
			break
		}

		fmt.Println(t.Format("2006/01/02 15:04:05"))
		fmt.Println(match[1:][1])
		fmt.Printf("%s\n\n", match[1:][0])

		resultSlice = append(resultSlice, match[1:][0:2])
	}
	if len(resultSlice) == 0 {
		return nil, errors.New("no records in the last 24 hours")
	}
	return resultSlice, nil
}
