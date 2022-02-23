package crawler

import (
	. "SecCrawler/config"
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
	if Cfg.Crawler.DongJian.Enabled {
		register.RegisterCrawler(&DongJian{})
	}
}
