package bot

import (
	"SecCrawler/config"
	"SecCrawler/register"
)

var cfg = config.GetGlobalConfig()

func init() {
	if cfg.DingBot.Enabled {
		register.RegisterBot(&DingBot{})
	}
	if cfg.FeishuBot.Enabled {
		register.RegisterBot(&FeishuBot{})
	}
	if cfg.HexQBot.Enabled {
		register.RegisterBot(&HexQBot{})
	}
	if cfg.ServerChan.Enabled {
		register.RegisterBot(&ServerChan{})
	}
	if cfg.WecomBot.Enabled {
		register.RegisterBot(&WecomBot{})
	}

}
