package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/robfig/cron"
)

func main() {
	if !cfg.Debug {
		_cron := cron.New()
		spec := fmt.Sprintf("0 0 %d * * ?", cfg.CronTime)
		err := _cron.AddFunc(spec, crawler)
		// err := _cron.AddFunc("0 */1 * * * ?", crawler) //每分钟
		if err != nil {
			log.Fatalf("add cron error: %s\n", err.Error())
		}

		_cron.Start()
		defer _cron.Stop()
		select {}
	} else {
		crawler()
	}
}

func crawler() {
	fmt.Printf("%s\n[♥︎] crawler start at %s\n%s\n\n", strings.Repeat("-", 47), currentTime(), strings.Repeat("-", 47))

	if cfg.EdgeForum.Enabled {
		var edgeForumResult [][]string
		var err error
		edgeForumResult, err = getEdgeForum()
		if err != nil {
			log.Printf("crawl [EdgeForum] error: %s\n\n", err.Error())
		} else {

			if cfg.WecomBot.Enabled {
				msg := wecomBotFormat(edgeForumResult, "EdgeForum")
				err := sendWecomBot(msg)
				if err != nil {
					log.Printf("send [EdgeForum] to [WecomBot] error: %s\n", err.Error())
				}
			}

			if cfg.HexQBot.Enabled {
				msg := commonFormat(edgeForumResult, "EdgeForum")
				err := sendHexQBot(msg)
				if err != nil {
					log.Printf("send [EdgeForum] to [HexQBot] error: %s\n", err.Error())
				}
			}

			if cfg.ServerChan.Enabled {
				title, msg := serverChanFormat(edgeForumResult, "EdgeForum")
				err := sendServerChan(title, msg)
				if err != nil {
					log.Printf("send [EdgeForum] to [ServerChan] error: %s\n", err.Error())
				}
			}

			if cfg.FeishuBot.Enabled {
				msg := commonFormat(edgeForumResult, "EdgeForum")
				err := sendFeishuBot(msg)
				if err != nil {
					log.Printf("send [EdgeForum] to [FeishuBot] error: %s\n", err.Error())
				}
			}

			if cfg.DingBot.Enabled {
				msg := commonFormat(edgeForumResult, "EdgeForum")
				err := sendDingBot(msg)
				if err != nil {
					log.Printf("send [EdgeForum] to [DingBot] error: %s\n", err.Error())
				}
			}

			// TODO: other bot

		}
	}

	if cfg.XianZhi.Enabled {
		var xianZhiResult [][]string
		var err error
		xianZhiResult, err = getXianZhi()
		if err != nil {
			log.Printf("crawl [XianZhi] error: %s\n\n", err.Error())
		} else {

			if cfg.WecomBot.Enabled {
				msg := wecomBotFormat(xianZhiResult, "XianZhi")
				err := sendWecomBot(msg)
				if err != nil {
					log.Printf("send [XianZhi] to [WecomBot] error: %s\n", err.Error())
				}
			}

			if cfg.HexQBot.Enabled {
				msg := commonFormat(xianZhiResult, "XianZhi")
				err := sendHexQBot(msg)
				if err != nil {
					log.Printf("send [XianZhi] to [HexQBot] error: %s\n", err.Error())
				}
			}

			if cfg.ServerChan.Enabled {
				title, msg := serverChanFormat(xianZhiResult, "XianZhi")
				err := sendServerChan(title, msg)
				if err != nil {
					log.Printf("send [XianZhi] to [ServerChan] error: %s\n", err.Error())
				}
			}

			if cfg.FeishuBot.Enabled {
				msg := commonFormat(xianZhiResult, "XianZhi")
				err := sendFeishuBot(msg)
				if err != nil {
					log.Printf("send [XianZhi] to [FeishuBot] error: %s\n", err.Error())
				}
			}

			if cfg.DingBot.Enabled {
				msg := commonFormat(xianZhiResult, "XianZhi")
				err := sendDingBot(msg)
				if err != nil {
					log.Printf("send [XianZhi] to [DingBot] error: %s\n", err.Error())
				}
			}

			// TODO: other bot

		}
	}

	if cfg.SeebugPaper.Enabled {
		var seebugPaperResult [][]string
		var err error
		seebugPaperResult, err = getSeebugPaper()
		if err != nil {
			log.Printf("crawl [SeebugPaper] error: %s\n\n", err.Error())
		} else {

			if cfg.WecomBot.Enabled {
				msg := wecomBotFormat(seebugPaperResult, "SeebugPaper")
				err := sendWecomBot(msg)
				if err != nil {
					log.Printf("send [SeebugPaper] to [WecomBot] error: %s\n", err.Error())
				}
			}

			if cfg.HexQBot.Enabled {
				msg := commonFormat(seebugPaperResult, "SeebugPaper")
				err := sendHexQBot(msg)
				if err != nil {
					log.Printf("send [SeebugPaper] to [HexQBot] error: %s\n", err.Error())
				}
			}

			if cfg.ServerChan.Enabled {
				title, msg := serverChanFormat(seebugPaperResult, "SeebugPaper")
				err := sendServerChan(title, msg)
				if err != nil {
					log.Printf("send [SeebugPaper] to [ServerChan] error: %s\n", err.Error())
				}
			}

			if cfg.FeishuBot.Enabled {
				msg := commonFormat(seebugPaperResult, "SeebugPaper")
				err := sendFeishuBot(msg)
				if err != nil {
					log.Printf("send [SeebugPaper] to [FeishuBot] error: %s\n", err.Error())
				}
			}

			if cfg.DingBot.Enabled {
				msg := commonFormat(seebugPaperResult, "SeebugPaper")
				err := sendDingBot(msg)
				if err != nil {
					log.Printf("send [SeebugPaper] to [DingBot] error: %s\n", err.Error())
				}
			}

			// TODO: other bot

		}
	}

	if cfg.Anquanke.Enabled {
		var anquankeResult [][]string
		var err error
		anquankeResult, err = getAnquanke()
		if err != nil {
			log.Printf("crawl [Anquanke] error: %s\n\n", err.Error())
		} else {

			if cfg.WecomBot.Enabled {
				msg := wecomBotFormat(anquankeResult, "Anquanke")
				err := sendWecomBot(msg)
				if err != nil {
					log.Printf("send [Anquanke] to [WecomBot] error: %s\n", err.Error())
				}
			}

			if cfg.HexQBot.Enabled {
				msg := commonFormat(anquankeResult, "Anquanke")
				err := sendHexQBot(msg)
				if err != nil {
					log.Printf("send [Anquanke] to [HexQBot] error: %s\n", err.Error())
				}
			}

			if cfg.ServerChan.Enabled {
				title, msg := serverChanFormat(anquankeResult, "Anquanke")
				err := sendServerChan(title, msg)
				if err != nil {
					log.Printf("send [Anquanke] to [ServerChan] error: %s\n", err.Error())
				}
			}

			if cfg.FeishuBot.Enabled {
				msg := commonFormat(anquankeResult, "Anquanke")
				err := sendFeishuBot(msg)
				if err != nil {
					log.Printf("send [Anquanke] to [FeishuBot] error: %s\n", err.Error())
				}
			}

			if cfg.DingBot.Enabled {
				msg := commonFormat(anquankeResult, "Anquanke")
				err := sendDingBot(msg)
				if err != nil {
					log.Printf("send [Anquanke] to [DingBot] error: %s\n", err.Error())
				}
			}

			// TODO: other bot

		}
	}
}
