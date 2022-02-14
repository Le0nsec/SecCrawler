package register

import "fmt"

type CrawlerConfig struct {
	Name        string // 站点名称
	Description string // 站点描述
}

type Crawler interface {
	Config() CrawlerConfig    // 爬虫爬取的站点名称与描述
	Get() ([][]string, error) // 爬虫爬取方法
}

var crawlerMap = map[string]Crawler{}

func RegisterCrawler(crawler Crawler) {
	fmt.Printf("[+] register crawler: [%s]\n", crawler.Config().Name)
	crawlerMap[crawler.Config().Name] = crawler
}

func GetCrawlerMap() map[string]Crawler {
	return crawlerMap
}

func GetCrawler(name string) (crawler Crawler, ok bool) {
	crawler, ok = crawlerMap[name]
	return
}
