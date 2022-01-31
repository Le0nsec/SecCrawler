package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// getEdgeForum 获取棱角社区前24小时内文章。
func getEdgeForum() ([][]string, error) {
	client := &http.Client{
		Timeout: time.Duration(4) * time.Second,
	}
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
	fmt.Printf("%s\n[*] [EdgeForum] crawler result:\n%s\n\n", strings.Repeat("-", 30), currentTime())
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

// getXianZhi 获取先知安全技术社区前24小时内文章。
func getXianZhi() ([][]string, error) {
	text, err := fetchXianZhiBySelenium()
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`<entry><title>.*?</title><link href=".*?" rel="alternate"></link><published>(.*?)</published><id>(.*?)</id><summary type="html">(.*?)</summary></entry>`)
	result := re.FindAllStringSubmatch(strings.TrimSpace(text), -1)

	var resultSlice [][]string
	fmt.Printf("%s\n[*] [XianZhi] crawler result:\n%s\n\n", strings.Repeat("-", 30), currentTime())
	for _, match := range result {
		t, err := time.Parse(time.RFC3339, match[1:][0])
		if err != nil {
			return nil, err
		}
		if !isIn24Hours(t) {
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
	_, err := selenium.NewChromeDriverService(cfg.ChromeDriver, 29515, opts...)
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

// getSeebugPaper 获取Paper Seebug（知道创宇）前24小时内文章。
func getSeebugPaper() ([][]string, error) {
	client := &http.Client{
		Timeout: time.Duration(4) * time.Second,
	}
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
	fmt.Printf("%s\n[*] [SeebugPaper] crawler result:\n%s\n\n", strings.Repeat("-", 30), currentTime())
	for _, match := range result {
		utc, _ := time.LoadLocation("UTC")
		t, err := time.ParseInLocation(time.RFC1123Z, match[1:][2], utc)
		if err != nil {
			return nil, err
		}

		time_zone := time.FixedZone("CST", 8*3600)
		if !isIn24Hours(t.In(time_zone)) {
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
		// slice中title和url调换位置，以符合统一的format
		for _, item := range resultSlice {
			item[0], item[1] = item[1], item[0]
		}
	}
	if len(resultSlice) == 0 {
		return nil, errors.New("no records in the last 24 hours")
	}
	return resultSlice, nil
}

// getAnquanke 获取安全客前24小时内文章。
func getAnquanke() ([][]string, error) {
	client := &http.Client{
		Timeout: time.Duration(4) * time.Second,
	}
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
	fmt.Printf("%s\n[*] [Anquanke] crawler result:\n%s\n\n", strings.Repeat("-", 30), currentTime())
	for _, match := range result {
		match[1:][0] = "https://www.anquanke.com" + match[1:][0]
		time_zone := time.FixedZone("CST", 8*3600)
		t, err := time.ParseInLocation("2006-01-02 15:04:05", match[1:][2], time_zone)
		if err != nil {
			return nil, err
		}

		if !isIn24Hours(t) {
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

// getTttang 获取跳跳糖前24小时内文章。
func getTttang() ([][]string, error) {
	client := &http.Client{
		Timeout: time.Duration(4) * time.Second,
	}
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(body)

	re := regexp.MustCompile(`<item><title>([\w\W]*?)</title><link>([\w\W]*?)</link><description>[\w\W]*?</description><dc:creator xmlns:dc="http://purl.org/dc/elements/1.1/">[\w\W]*?</dc:creator><pubDate>([\w\W]*?)</pubDate><guid>[\w\W]*?</guid></item>`)
	result := re.FindAllStringSubmatch(strings.TrimSpace(bodyString), -1)

	var resultSlice [][]string
	fmt.Printf("%s\n[*] [Tttang] crawler result:\n%s\n\n", strings.Repeat("-", 30), currentTime())
	for _, match := range result {
		time_zone := time.FixedZone("CST", 8*3600)
		t, err := time.ParseInLocation(time.RFC1123Z, match[1:][2], time_zone)
		if err != nil {
			return nil, err
		}

		if !isIn24Hours(t.In(time_zone)) {
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
		// slice中title和url调换位置，以符合统一的format
		for _, item := range resultSlice {
			item[0], item[1] = item[1], item[0]
		}
	}
	if len(resultSlice) == 0 {
		return nil, errors.New("no records in the last 24 hours")
	}
	return resultSlice, nil
}

func getQiAnXin() ([][]string, error) {
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
	fmt.Printf("%s\n[*] [QiAnXin] crawler result:\n%s\n\n", strings.Repeat("-", 30), currentTime())
	for _, match := range result {
		time_zone := time.FixedZone("CST", 8*3600)
		t, err := time.ParseInLocation("2006-01-02 15:04:05", match[1:][2], time_zone)
		if err != nil {
			return nil, err
		}

		if !isIn24Hours(t.In(time_zone)) {
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
