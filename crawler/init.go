package crawler

import (
	. "SecCrawler/config"
	"SecCrawler/crawler/lab"
	"SecCrawler/register"
)

func CrawlerInit() {
	if Cfg.Crawler.Anquanke.Enabled {
		register.RegisterCrawler(&Anquanke{})
	}
	if Cfg.Crawler.EdgeForum.Enabled {
		register.RegisterCrawler(&EdgeForum{})
	}
	if Cfg.Crawler.QiAnXin.Enabled {
		register.RegisterCrawler(&QiAnXin{})
	}
	if Cfg.Crawler.SeebugPaper.Enabled {
		register.RegisterCrawler(&SeebugPaper{})
	}
	if Cfg.Crawler.Tttang.Enabled {
		register.RegisterCrawler(&Tttang{})
	}
	if Cfg.Crawler.XianZhi.Enabled {
		register.RegisterCrawler(&XianZhi{})
	}
	// 暂时删除洞见微信聚合，优化推送体验
	// if Cfg.Crawler.DongJian.Enabled {
	// 	register.RegisterCrawler(&DongJian{})
	// }
	if Cfg.Crawler.Lab.Enabled {
		register.RegisterCrawler(&lab.Lab{})
	}
	if Cfg.Crawler.HuoxianZone.Enabled {
		register.RegisterCrawler(&HuoxianZone{})
	}
}
