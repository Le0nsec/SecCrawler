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
)

type EdgeForum struct{}

func (crawler EdgeForum) Config() register.CrawlerConfig {
	return register.CrawlerConfig{
		Name:        "EdgeForum",
		Description: "棱角社区攻防日报",
	}
}

// Get 获取棱角社区前24小时内文章。
func (crawler EdgeForum) Get() ([][]string, error) {
	client := utils.CrawlerClient()

	req, err := http.NewRequest("GET", "https://forum.ywhack.com/forumdisplay.php?fid=59&orderby=lastpost&filter=86400", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
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

	//去除连续的换行符
	re, _ = regexp.Compile(`\s{2,}`)
	bodyString = re.ReplaceAllString(bodyString, "")
	// fmt.Println(bodyString)

	re = regexp.MustCompile(`<a href="(.*?)" target="_blank">(.*?)</a>.*?<img src="" style="vertical-align: top;margin-top: 2px;"></div><small class="card-subtitle text-muted">.*?<span class="badge badge-tag">.*?</span></small>`)
	result := re.FindAllStringSubmatch(strings.TrimSpace(bodyString), -1)
	// fmt.Println(result)

	var resultSlice [][]string
	fmt.Printf("[*] [EdgeForum] crawler result:\n%s\n\n", utils.CurrentTime())
	for _, match := range result {
		fmt.Printf("%s\n", match[1:][1])
		fmt.Printf("%s\n\n", match[1:][0])
		resultSlice = append(resultSlice, match[1:])
	}
	if len(resultSlice) == 0 {
		return nil, errors.New("no records in the last 24 hours")
	}
	return resultSlice, nil
}
