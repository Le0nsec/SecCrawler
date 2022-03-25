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
		if Cfg.Crawler.Lab.NoahLab.Enabled {
			register.RegisterCrawler(&lab.NoahLab{})
		}
		if Cfg.Crawler.Lab.Blog360.Enabled {
			register.RegisterCrawler(&lab.Blog360{})
		}
		if Cfg.Crawler.Lab.Nsfocus.Enabled {
			register.RegisterCrawler(&lab.Nsfocus{})
		}
		if Cfg.Crawler.Lab.Xlab.Enabled {
			register.RegisterCrawler(&lab.Xlab{})
		}
		if Cfg.Crawler.Lab.AlphaLab.Enabled {
			register.RegisterCrawler(&lab.AlphaLab{})
		}
		if Cfg.Crawler.Lab.Netlab.Enabled {
			register.RegisterCrawler(&lab.Netlab{})
		}
		if Cfg.Crawler.Lab.RiskivyBlog.Enabled {
			register.RegisterCrawler(&lab.RiskivyBlog{})
		}
		if Cfg.Crawler.Lab.TSRCBlog.Enabled {
			register.RegisterCrawler(&lab.TSRCBlog{})
		}
		if Cfg.Crawler.Lab.X1cT34m.Enabled {
			register.RegisterCrawler(&lab.X1cT34m{})
		}
	}
}
