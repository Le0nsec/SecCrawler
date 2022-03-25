package lab

import (
	. "SecCrawler/config"
	"SecCrawler/register"
	"errors"
	"log"
)

type Lab struct{}

func (crawler Lab) Config() register.CrawlerConfig {
	return register.CrawlerConfig{
		Name:        "Lab",
		Description: "实验室文章",
	}
}

// Get 获取 Lab 前24小时内文章。
func (crawler Lab) Get() ([][]string, error) {
	var resultSlice [][]string

	if Cfg.Crawler.Lab.NoahLab.Enabled {
		resultSlice = tmpCrawler(resultSlice, NoahLab{})
	}
	if Cfg.Crawler.Lab.Blog360.Enabled {
		resultSlice = tmpCrawler(resultSlice, Blog360{})
	}
	if Cfg.Crawler.Lab.Nsfocus.Enabled {
		resultSlice = tmpCrawler(resultSlice, Nsfocus{})
	}
	if Cfg.Crawler.Lab.Xlab.Enabled {
		resultSlice = tmpCrawler(resultSlice, Xlab{})
	}
	if Cfg.Crawler.Lab.AlphaLab.Enabled {
		resultSlice = tmpCrawler(resultSlice, AlphaLab{})
	}
	if Cfg.Crawler.Lab.Netlab.Enabled {
		resultSlice = tmpCrawler(resultSlice, Netlab{})
	}
	if Cfg.Crawler.Lab.RiskivyBlog.Enabled {
		resultSlice = tmpCrawler(resultSlice, RiskivyBlog{})
	}
	if Cfg.Crawler.Lab.TSRCBlog.Enabled {
		resultSlice = tmpCrawler(resultSlice, TSRCBlog{})
	}
	if Cfg.Crawler.Lab.X1cT34m.Enabled {
		resultSlice = tmpCrawler(resultSlice, X1cT34m{})
	}

	if len(resultSlice) == 0 {
		return nil, errors.New("no records in the last 24 hours")
	}
	return resultSlice, nil
}

func tmpCrawler(s [][]string, crawler register.Crawler) [][]string {
	crawlerResult, err := crawler.Get()
	if err != nil {
		log.Printf("crawl [%s] error: %s\n\n", crawler.Config().Name, err.Error())
	}
	s = append(s, crawlerResult...)
	return s
}
