package config

type Config struct {
	ChromeDriver string `yaml:"ChromeDriver"`

	Proxy   ProxyStruct   `yaml:"Proxy"`
	Cron    CronStruct    `yaml:"Cron"`
	Api     ApiStruct     `yaml:"Api"`
	Crawler CrawlerStruct `yaml:"Crawler"`
	Bot     BotStruct     `yaml:"Bot"`
}

type CronStruct struct {
	Enabled bool  `yaml:"enabled"`
	Time    uint8 `yaml:"time"`
}

type ProxyStruct struct {
	ProxyUrl            string `yaml:"ProxyUrl"`
	CrawlerProxyEnabled bool   `yaml:"CrawlerProxyEnabled"`
	BotProxyEnabled     bool   `yaml:"BotProxyEnabled"`
}

type ApiStruct struct {
	Enabled bool   `yaml:"enabled"`
	Debug   bool   `yaml:"debug"`
	Host    string `yaml:"host"`
	Port    uint16 `yaml:"port"`
	Auth    string `yaml:"auth"`
}

type CrawlerStruct struct {
	EdgeForum   EdgeForumStruct   `yaml:"EdgeForum"`
	XianZhi     XianZhiStruct     `yaml:"XianZhi"`
	SeebugPaper SeebugPaperStruct `yaml:"SeebugPaper"`
	Anquanke    AnquankeStruct    `yaml:"Anquanke"`
	Tttang      TttangStruct      `yaml:"Tttang"`
	QiAnXin     QiAnXinStruct     `yaml:"QiAnXin"`
	// DongJian    DongJianStruct    `yaml:"DongJian"`
	Lab LabStruct `yaml:"Lab"`
}

type BotStruct struct {
	WecomBot   WecomBotStruct   `yaml:"WecomBot"`
	FeishuBot  FeishuBotStruct  `yaml:"FeishuBot"`
	DingBot    DingBotStruct    `yaml:"DingBot"`
	HexQBot    HexQBotStruct    `yaml:"HexQBot"`
	ServerChan ServerChanStruct `yaml:"ServerChan"`
	WgpSecBot  WgpSecBotStruct  `yaml:"WgpSecBot"`
	CqHttpBot  CqHttpBotStruct  `yaml:"CqHttpBot"`
}

type WecomBotStruct struct {
	Enabled bool   `yaml:"enabled"`
	Key     string `yaml:"key"`
	Timeout uint8  `yaml:"timeout"`
}

type FeishuBotStruct struct {
	Enabled bool   `yaml:"enabled"`
	Key     string `yaml:"key"`
	Timeout uint8  `yaml:"timeout"`
}

type DingBotStruct struct {
	Enabled bool   `yaml:"enabled"`
	Token   string `yaml:"token"`
	Timeout uint8  `yaml:"timeout"`
}

type HexQBotStruct struct {
	Enabled bool   `yaml:"enabled"`
	Api     string `yaml:"api"`
	QQGroup uint64 `yaml:"qqgroup"`
	Key     string `yaml:"key"`
	Timeout uint8  `yaml:"timeout"`
}

type ServerChanStruct struct {
	Enabled bool   `yaml:"enabled"`
	SendKey string `yaml:"sendkey"`
	Timeout uint8  `yaml:"timeout"`
}

type WgpSecBotStruct struct {
	Enabled bool   `yaml:"enabled"`
	Key     string `yaml:"key"`
	Timeout uint8  `yaml:"timeout"`
}

type CqHttpBotStruct struct {
	Enabled bool   `yaml:"enabled"`
	Api     string `yaml:"api"`
	QQGroup uint64 `yaml:"qqgroup"`
	Key     string `yaml:"key"`
	Timeout uint8  `yaml:"timeout"`
}

type EdgeForumStruct struct {
	Enabled bool `yaml:"enabled"`
}

type XianZhiStruct struct {
	Enabled         bool   `yaml:"enabled"`
	UseChromeDriver bool   `yaml:"UseChromeDriver"`
	CustomRSSURL    string `yaml:"CustomRSSURL"`
}

type SeebugPaperStruct struct {
	Enabled bool `yaml:"enabled"`
}

type AnquankeStruct struct {
	Enabled bool `yaml:"enabled"`
}

type TttangStruct struct {
	Enabled bool `yaml:"enabled"`
}

type QiAnXinStruct struct {
	Enabled bool `yaml:"enabled"`
}

type DongJianStruct struct {
	Enabled bool `yaml:"enabled"`
}

type LabStruct struct {
	Enabled     bool              `yaml:"enabled"`
	NoahLab     NoahLabStruct     `yaml:"NoahLab"`
	Blog360     Blog360Struct     `yaml:"Blog360"`
	Nsfocus     NsfocusStruct     `yaml:"Nsfocus"`
	Xlab        XlabStruct        `yaml:"Xlab"`
	AlphaLab    AlphaLabStruct    `yaml:"AlphaLab"`
	Netlab      NetlabStruct      `yaml:"Netlab"`
	RiskivyBlog RiskivyBlogStruct `yaml:"RiskivyBlog"`
	TSRCBlog    TSRCBlogStruct    `yaml:"TSRCBlog"`
	X1cT34m     X1cT34mStruct     `yaml:"X1cT34m"`
}

type NoahLabStruct struct {
	Enabled bool `yaml:"enabled"`
}
type Blog360Struct struct {
	Enabled bool `yaml:"enabled"`
}
type NsfocusStruct struct {
	Enabled bool `yaml:"enabled"`
}
type XlabStruct struct {
	Enabled bool `yaml:"enabled"`
}
type AlphaLabStruct struct {
	Enabled bool `yaml:"enabled"`
}
type NetlabStruct struct {
	Enabled bool `yaml:"enabled"`
}
type RiskivyBlogStruct struct {
	Enabled bool `yaml:"enabled"`
}
type TSRCBlogStruct struct {
	Enabled bool `yaml:"enabled"`
}
type X1cT34mStruct struct {
	Enabled bool `yaml:"enabled"`
}
