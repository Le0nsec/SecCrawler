package crawler

import (
	"SecCrawler/config"
	"SecCrawler/register"
)

var cfg = config.GetGlobalConfig()

func init() {
	if cfg.Anquanke.Enabled {
		register.RegisterCrawler(&Anquanke{})
	}
	if cfg.EdgeForum.Enabled {
		register.RegisterCrawler(&EdgeForum{})
	}
	if cfg.QiAnXin.Enabled {
		register.RegisterCrawler(&QiAnXin{})
	}
	if cfg.SeebugPaper.Enabled {
		register.RegisterCrawler(&SeebugPaper{})
	}
	if cfg.Tttang.Enabled {
		register.RegisterCrawler(&Tttang{})
	}
	if cfg.XianZhi.Enabled {
		register.RegisterCrawler(&XianZhi{})
	}
}
