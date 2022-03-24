package crawler

import (
	"SecCrawler/config"
	. "SecCrawler/config"
	"SecCrawler/register"
	"SecCrawler/utils"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
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
	var resultSlice [][]string

	if config.Cfg.Crawler.XianZhi.UseChromeDriver {
		text, err := fetchXianZhiBySelenium()
		if err != nil {
			return nil, err
		}

		re := regexp.MustCompile(`<entry><title>.*?</title><link href=".*?" rel="alternate"></link><published>(.*?)</published><id>(.*?)</id><summary type="html">(.*?)</summary></entry>`)
		result := re.FindAllStringSubmatch(strings.TrimSpace(text), -1)

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
	} else {
		if config.Cfg.Crawler.XianZhi.CustomRSSURL == "" {
			return nil, errors.New("ChromeDriver is disabled and no custom RSS URL is specified")
		}
		client := utils.CrawlerClient()

		req, err := http.NewRequest("GET", config.Cfg.Crawler.XianZhi.CustomRSSURL, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
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

		fmt.Printf("[*] [XianZhi] crawler result:\n%s\n\n", utils.CurrentTime())

		for _, item := range feed.Items {
			t, err := time.Parse(time.RFC3339, item.Published)
			if err != nil {
				return nil, err
			}
			if !utils.IsIn24Hours(t) {
				break
			}

			fmt.Println(t.Format("2006/01/02 15:04:05"))
			fmt.Println(item.Title)
			fmt.Printf("%s\n\n", item.Link)

			var s []string
			s = append(s, item.Link, item.Title)
			resultSlice = append(resultSlice, s)
		}
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

	args := []string{
		"--headless",
		"--no-sandbox",
		"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36",
	}
	if config.Cfg.Proxy.CrawlerProxyEnabled {
		// 设置代理
		proxyArgs := fmt.Sprintf("--proxy-server=%s", config.Cfg.Proxy.ProxyUrl)
		args = append(args, proxyArgs)
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args:  args,
	}

	caps.AddChrome(chromeCaps)
	service, err := selenium.NewChromeDriverService(Cfg.ChromeDriver, 29515, opts...)
	if err != nil {
		return "", err
	}
	defer service.Stop()

	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 29515))
	if err != nil {
		return "", err
	}
	defer webDriver.Quit()

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

	time.Sleep(3 * time.Second)
	return text, nil
}
