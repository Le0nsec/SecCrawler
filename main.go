package main

import (
	"SecCrawler/config"
	"SecCrawler/crawler"
	"SecCrawler/send"
	"SecCrawler/untils"
	"fmt"
	"github.com/robfig/cron"
	"log"
)

func main() {
	//a := "sssssssss"
	//send.SendDingBot(a)
	//spider()
	_cron := cron.New()
	spec := fmt.Sprintf("0 0 %d * * ?", config.GetConfig().CronTime)
	fmt.Println(spec)
	err := _cron.AddFunc(spec, spider)
	// err := _cron.AddFunc("0 */1 * * * ?", spider) //每分钟
	if err != nil {
		log.Fatalf("add cron error: %s\n", err.Error())
	}

	_cron.Start()
	defer _cron.Stop()
	select {}
}

func spider() {

	fmt.Printf("-----------------------------------------------\n[♥︎] spider start at %s\n-----------------------------------------------\n\n", untils.CurrentTime())

	if config.GetConfig().EdgeForum.Enabled {
		var edgeForumResult [][]string
		var err error
		edgeForumResult, err = crawler.GetEdgeForum()
		if err != nil {
			log.Printf("crawl [EdgeForum] error: %s\n\n", err.Error())
		} else {

			if config.Cfg.WecomBot.Enabled {
				msg := untils.WecomBotFormat(edgeForumResult, "EdgeForum")
				err := send.SendWecomBot(msg)
				if err != nil {
					log.Printf("send [EdgeForum] to [WecomBot] error: %s\n", err.Error())
				}
			}

			if config.Cfg.HexQBot.Enabled {
				msg := untils.CommonFormat(edgeForumResult, "EdgeForum")
				err := send.SendHexQBot(msg)
				if err != nil {
					log.Printf("send [EdgeForum] to [HexQBot] error: %s\n", err.Error())
				}
			}

			if config.Cfg.ServerChan.Enabled {
				title, msg := untils.ServerChanFormat(edgeForumResult, "EdgeForum")
				err := send.SendServerChan(title, msg)
				if err != nil {
					log.Printf("send [EdgeForum] to [ServerChan] error: %s\n", err.Error())
				}
			}

			if config.Cfg.FeishuBot.Enabled {
				msg := untils.CommonFormat(edgeForumResult, "EdgeForum")
				err := send.SendFeishuBot(msg)
				if err != nil {
					log.Printf("send [EdgeForum] to [FeishuBot] error: %s\n", err.Error())
				}
			}

			fmt.Println(config.Cfg.DingBot.Enabled)
			fmt.Println(config.Cfg.DingBot.Token)
			if config.Cfg.DingBot.Enabled {
				fmt.Println("Ddd")
				msg := untils.CommonFormat(edgeForumResult, "EdgeForum")
				err := send.SendDingBot(msg)
				if err != nil {
					log.Printf("send [EdgeForum] to [DingBot] error: %s\n", err.Error())
				}
			}

			// TODO: other bot

		}
	}

	if config.Cfg.XianZhi.Enabled {
		var xianZhiResult [][]string
		var err error
		xianZhiResult, err = crawler.GetXianZhi()
		if err != nil {
			log.Printf("crawl [XianZhi] error: %s\n\n", err.Error())
		} else {

			if config.Cfg.WecomBot.Enabled {
				msg := untils.WecomBotFormat(xianZhiResult, "XianZhi")
				err := send.SendWecomBot(msg)
				if err != nil {
					log.Printf("send [XianZhi] to [WecomBot] error: %s\n", err.Error())
				}
			}

			if config.Cfg.HexQBot.Enabled {
				msg := untils.CommonFormat(xianZhiResult, "XianZhi")
				err := send.SendHexQBot(msg)
				if err != nil {
					log.Printf("send [XianZhi] to [HexQBot] error: %s\n", err.Error())
				}
			}

			if config.Cfg.ServerChan.Enabled {
				title, msg := untils.ServerChanFormat(xianZhiResult, "XianZhi")
				err := send.SendServerChan(title, msg)
				if err != nil {
					log.Printf("send [XianZhi] to [ServerChan] error: %s\n", err.Error())
				}
			}

			if config.Cfg.FeishuBot.Enabled {
				msg := untils.CommonFormat(xianZhiResult, "XianZhi")
				err := send.SendFeishuBot(msg)
				if err != nil {
					log.Printf("send [XianZhi] to [FeishuBot] error: %s\n", err.Error())
				}
			}

			if config.Cfg.DingBot.Enabled {
				msg := untils.CommonFormat(xianZhiResult, "XianZhi")
				err := send.SendDingBot(msg)
				if err != nil {
					log.Printf("send [XianZhi] to [DingBot] error: %s\n", err.Error())
				}
			}

			// TODO: other bot

		}
	}

	if config.Cfg.SeebugPaper.Enabled {
		var seebugPaperResult [][]string
		var err error
		seebugPaperResult, err = crawler.GetSeebugPaper()
		if err != nil {
			log.Printf("crawl [SeebugPaper] error: %s\n\n", err.Error())
		} else {

			if config.Cfg.WecomBot.Enabled {
				msg := untils.WecomBotFormat(seebugPaperResult, "SeebugPaper")
				err := send.SendWecomBot(msg)
				if err != nil {
					log.Printf("send [SeebugPaper] to [WecomBot] error: %s\n", err.Error())
				}
			}

			if config.Cfg.HexQBot.Enabled {
				msg := untils.CommonFormat(seebugPaperResult, "SeebugPaper")
				err := send.SendHexQBot(msg)
				if err != nil {
					log.Printf("send [SeebugPaper] to [HexQBot] error: %s\n", err.Error())
				}
			}

			fmt.Println(config.Cfg.ServerChan.Enabled)
			if false{
				fmt.Println("dd")
			}
			if config.Cfg.ServerChan.Enabled {
				title, msg := untils.ServerChanFormat(seebugPaperResult, "SeebugPaper")
				err := send.SendServerChan(title, msg)
				if err != nil {
					log.Printf("send [SeebugPaper] to [ServerChan] error: %s\n", err.Error())
				}
			}

			if config.Cfg.FeishuBot.Enabled {
				msg := untils.CommonFormat(seebugPaperResult, "SeebugPaper")
				err := send.SendFeishuBot(msg)
				if err != nil {
					log.Printf("send [SeebugPaper] to [FeishuBot] error: %s\n", err.Error())
				}
			}

			fmt.Println("DingBot",config.Cfg.DingBot.Enabled)
			if config.Cfg.DingBot.Enabled {
				msg := untils.CommonFormat(seebugPaperResult, "SeebugPaper")
				err := send.SendDingBot(msg)
				if err != nil {
					log.Printf("send [SeebugPaper] to [DingBot] error: %s\n", err.Error())
				}
			}

			// TODO: other bot

		}
	}

	if config.Cfg.Anquanke.Enabled {
		var anquankeResult [][]string
		var err error
		anquankeResult, err = crawler.GetAnquanke()
		if err != nil {
			log.Printf("crawl [Anquanke] error: %s\n\n", err.Error())
		} else {

			if config.Cfg.WecomBot.Enabled {
				msg := untils.WecomBotFormat(anquankeResult, "Anquanke")
				err := send.SendWecomBot(msg)
				if err != nil {
					log.Printf("send [Anquanke] to [WecomBot] error: %s\n", err.Error())
				}
			}

			if config.Cfg.HexQBot.Enabled {
				msg := untils.CommonFormat(anquankeResult, "Anquanke")
				err := send.SendHexQBot(msg)
				if err != nil {
					log.Printf("send [Anquanke] to [HexQBot] error: %s\n", err.Error())
				}
			}

			if config.Cfg.ServerChan.Enabled {
				title, msg := untils.ServerChanFormat(anquankeResult, "Anquanke")
				err := send.SendServerChan(title, msg)
				if err != nil {
					log.Printf("send [Anquanke] to [ServerChan] error: %s\n", err.Error())
				}
			}

			if config.Cfg.FeishuBot.Enabled {
				msg := untils.CommonFormat(anquankeResult, "Anquanke")
				err := send.SendFeishuBot(msg)
				if err != nil {
					log.Printf("send [Anquanke] to [FeishuBot] error: %s\n", err.Error())
				}
			}



			if config.Cfg.DingBot.Enabled {
				msg := untils.CommonFormat(anquankeResult, "Anquanke")
				err := send.SendDingBot(msg)
				if err != nil {
					log.Printf("send [Anquanke] to [DingBot] error: %s\n", err.Error())
				}
			}

			// TODO: other bot

		}
	}
}
