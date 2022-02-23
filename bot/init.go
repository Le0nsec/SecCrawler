package bot

import (
	. "SecCrawler/config"
	"SecCrawler/register"
)

func BotInit() {
	if Cfg.Bot.DingBot.Enabled {
		register.RegisterBot(&DingBot{})
	}
	if Cfg.Bot.FeishuBot.Enabled {
		register.RegisterBot(&FeishuBot{})
	}
	if Cfg.Bot.HexQBot.Enabled {
		register.RegisterBot(&HexQBot{})
	}
	if Cfg.Bot.ServerChan.Enabled {
		register.RegisterBot(&ServerChan{})
	}
	if Cfg.Bot.WecomBot.Enabled {
		register.RegisterBot(&WecomBot{})
	}
	if Cfg.Bot.WgpSecBot.Enabled {
		register.RegisterBot(&WgpSecBot{})
	}

}
