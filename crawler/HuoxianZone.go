package crawler

import (
	"SecCrawler/register"
	"SecCrawler/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type HuoxianZone struct{}

type respDataJson struct {
	Data []struct {
		Attributes struct {
			Title        string `json:"title"`
			Slug         string `json:"slug"`
			LastPostedAt string `json:"lastPostedAt"`
		} `json:"attributes"`
		Relationships struct {
			Tags struct {
				Data []struct {
					Type string `json:"type"`
					Id   string `json:"id"`
				} `json:"data"`
			} `json:"tags"`
		} `json:"relationships"`
	} `json:"data"`
}

func (crawler HuoxianZone) Config() register.CrawlerConfig {
	return register.CrawlerConfig{
		Name:        "火线Zone",
		Description: "火线 Zone-安全攻防社区",
	}
}

func query20Data(page int) (*respDataJson, error) {

	client := utils.CrawlerClient()

	// 默认查询20条，从0开始。
	getUrl := "https://zone.huoxian.cn/api/discussions?include=tags&sort=-createdAt&page[offset]=" + strconv.Itoa(page)

	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("authority", "zone.huoxian.cn")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://zone.huoxian.cn/?sort=newest")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="119", "Chromium";v="119", "Not?A_Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	currentDataJson := &respDataJson{}
	err = json.NewDecoder(resp.Body).Decode(currentDataJson)
	if err != nil {
		return nil, err
	}

	return currentDataJson, nil

}

// Get 获取安全客前24小时内文章。
func (crawler HuoxianZone) Get() ([][]string, error) {

	var resultSlice [][]string
	zoneFetchDone := false

	for page := 0; ; page += 20 {

		ttt, err := query20Data(page)
		if err != nil {
			return nil, err
		}
		pageResult := *ttt

		fmt.Printf("[*] [HuoxianZone] crawler result:\n%s\n\n", utils.CurrentTime())

		for _, match := range pageResult.Data {

			isPaper := false
			for _, data := range match.Relationships.Tags.Data {
				if data.Type != "tags" {
					continue
				}
				// 官方公告，直接跳过，同时避免置顶文章影响时间判断
				if data.Id == "2" {
					break
				}
				if data.Id == "4" {
					isPaper = true
				}
			}
			if !isPaper {
				continue
			}

			timeZone := time.FixedZone("CST", 8*3600)
			t, err := time.ParseInLocation("2006-01-02T15:04:05+00:00", match.Attributes.LastPostedAt, timeZone)
			if err != nil {
				continue
			}

			if !utils.IsIn24Hours(t) {
				// 默认时间顺序是从近到远
				// 已经获取了所有24小时内的文章
				zoneFetchDone = true
				break
			}

			fmt.Println(t.Format("2006/01/02 15:04:05"))
			fmt.Println(match.Attributes.Title)
			paperUrl := "https://zone.huoxian.cn/d/" + match.Attributes.Slug
			fmt.Printf("%s\n\n", paperUrl)

			var s []string
			s = append(s, paperUrl, match.Attributes.Title)
			resultSlice = append(resultSlice, s)
		}

		if zoneFetchDone {
			break
		}

	}

	if len(resultSlice) == 0 {
		return nil, errors.New("no records in the last 24 hours")
	}
	return resultSlice, nil

}
