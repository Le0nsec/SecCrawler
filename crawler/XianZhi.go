package crawler

import (
	. "SecCrawler/config"
	"SecCrawler/register"
	"SecCrawler/utils"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type XianZhi struct{}

func (crawler XianZhi) Config() register.CrawlerConfig {
	return register.CrawlerConfig{
		Name:        "XianZhi",
		Description: "先知安全技术社区",
	}
}

// Get 获取先知安全技术社区前24小时内文章。
func (crawler XianZhi) Get() ([][]string, error) {
	text, err := fetchXianZhiBySelenium()
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`<entry><title>.*?</title><link href=".*?" rel="alternate"></link><published>(.*?)</published><id>(.*?)</id><summary type="html">(.*?)</summary></entry>`)
	result := re.FindAllStringSubmatch(strings.TrimSpace(text), -1)

	var resultSlice [][]string
	fmt.Printf("[*] [XianZhi] crawler result:\n%s\n\n", utils.CurrentTime())
	for _, match := range result {
		t, err := time.Parse(time.RFC3339, match[1:][0])
		if err != nil {
			return nil, err
		}
		if !utils.IsIn24Hours(t) {
			// 默认时间顺序是从近到远
			break
		}

		fmt.Println(t.Format("2006/01/02 15:04:05"))
		fmt.Println(match[1:][2])
		fmt.Printf("%s\n\n", match[1:][1])

		resultSlice = append(resultSlice, match[1:][1:])
	}
	if len(resultSlice) == 0 {
		return nil, errors.New("no records in the last 24 hours")
	}
	return resultSlice, nil
}

// fetchXianZhiBySelenium 使用selenium爬取先知社区，因为先知有cookie动态反爬。
func fetchXianZhiBySelenium() (string, error) {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless",
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36",
		},
	}

	caps.AddChrome(chromeCaps)
	_, err := selenium.NewChromeDriverService(Cfg.ChromeDriver, 29515, opts...)
	if err != nil {
		return "", err
	}

	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 29515))
	if err != nil {
		return "", err
	}

	// webDriver.AddCookie(&selenium.Cookie{
	// 	Name:  "defaultJumpDomain",
	// 	Value: "www",
	// })

	err = webDriver.Get("https://xz.aliyun.com/feed")
	if err != nil {
		return "", err
	}
	element, err := webDriver.FindElement(selenium.ByXPATH, "/html/body/pre")
	if err != nil {
		return "", err
	}
	text, err := element.Text()
	if err != nil {
		return "", err
	}
	return text, nil
}
