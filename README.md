<h1 align="center">
SecCrawler
</h1>


<h4 align="center">
一个方便安全研究人员获取每日安全日报的爬虫和推送程序，目前爬取范围包括先知社区、安全客、Seebug Paper、棱角社区，持续更新中。
</h4>


<p align="center">
  <a href="https://github.com/Le0nsec/SecCrawler/issues">
    <img src="https://img.shields.io/github/issues/Le0nsec/SecCrawler?style=flat-square">
  </a>
  <a href="https://github.com/Le0nsec/SecCrawler/network/members">
    <img src="https://img.shields.io/github/forks/Le0nsec/SecCrawler?style=flat-square">
  </a>
  <a href="https://github.com/Le0nsec/SecCrawler/stargazers">
    <img src="https://img.shields.io/github/stars/Le0nsec/SecCrawler?style=flat-square">
  </a>
  <a href="https://github.com/Le0nsec/SecCrawler/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/Le0nsec/SecCrawler?style=flat-square">
  </a>
  <a href="https://github.com/RichardLitt/standard-readme">
    <img src="https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square">
  </a>
  <a href="https://github.com/Le0nsec/SecCrawler/releases">
    <img src="https://img.shields.io/github/v/release/Le0nsec/SecCrawler?include_prereleases&style=flat-square">
  </a>
</p>



## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Install](#install)
- [Config](#config)
- [Contributing](#contributing)
- [License](#license)


## Introduction

SecCrawler 是一个跨平台的方便安全研究人员获取每日安全日报的爬虫和机器人推送程序，目前爬取范围包括先知社区、安全客、Seebug Paper、棱角社区，持续更新中。


程序使用yml格式的配置文件，第一次运行时会在当前文件夹自动生成`config.yml`配置文件模板，在配置文件中设置爬取的网站和推送机器人相关配置，目前包括在内的网站和推送的机器人在[Features](#features)中可以查看，可以设置每日推送的整点时间。


程序使用定时任务每天根据设置好的时间整点自动运行，编辑好相关配置后后台运行即可，示例运行命令：


```sh
$ nohup ./SecCrawler >> run.log 2>&1 &
# 或者使用screen
$ screen ./SecCrawler
$ ctrl a+d / control a+d # 回到主会话
```


注：由于在爬取先知安全社区时程序使用了 Selenium，用户需要手动下载`ChromeDriver`和`Chrome`浏览器。


ChromeDriver镜像站：http://npm.taobao.org/mirrors/chromedriver/


- Windows和Mac用户在[下载Chrome](https://www.google.cn/chrome/)并安装后，下载对应chrome版本的ChromeDriver并在配置文件`config.yml`中指定ChromeDriver的路径
- Linux用户在下载Chrome（链接如下）并安装后，同上编辑配置文件
    - [Debian/Ubuntu(64位.deb)](https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb)
    - [Fedora/openSUSE(64位.rpm)](https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm)


> Chrome浏览器可以访问`chrome://version/`查看版本
> 命令行可以使用`google-chrome-stable --version`查看版本


程序旨在帮助安全研究者自动化获取每日更新的安全文章，适用于每日安全日报推送，爬取的安全社区网站范围和支持推送的机器人持续增加中，欢迎在[issues](https://github.com/Le0nsec/SecCrawler/issues)中提供宝贵的建议。


:rocket: 目前 SecCrawler 已在MacOS Apple silicon 、Ubuntu 20.04运行测试通过。

## Features

- 爬取网站列表
    - [x] [先知安全社区](https://xz.aliyun.com/)
    - [x] [安全客](https://www.anquanke.com/knowledge) (安全知识专区)
    - [x] [Seebug Paper](https://paper.seebug.org/)
    - [x] [棱角安全社区](https://forum.ywhack.com/forum-59-1.html)
- 推送机器人列表
    - [x] [企业微信群机器人](https://work.weixin.qq.com/api/doc/90000/90136/91770)
    - [x] [HexQBot](https://github.com/Am473ur/HexQBot) (QQ群机器人 自建)
    - [x] [Server酱](https://sct.ftqq.com/)
    - [ ] [飞书群机器人](https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN)
    - [ ] [钉钉群机器人](https://open.dingtalk.com/document/robots/custom-robot-access)

## Install

你可以在[Releases](https://github.com/Le0nsec/SecCrawler-dev/releases)下载最新的SecCrawler。


或者从源码编译：


```sh
$ git clone https://github.com/Le0nsec/SecCrawler.git
$ cd SecCrawler
$ go build .
```


## Config
`config.yml`配置文件模板：

```yml

############### CronSetting ###############

# 设置每天整点爬取推送时间，范围 0 ~ 23（整数）
CronTime: 10
# 设置Selenium使用的ChromeDriver路径，支持相对路径或绝对路径（如果不爬取先知社区可以不用设置）
ChromeDriver: ./chromedriver/linux64


############### BotConfig ###############

# 企业微信群机器人
# https://work.weixin.qq.com/api/doc/90000/90136/91770
WecomBot:
  enabled: false
  key: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
  timeout: 2  # second

# HexQBot
# https://github.com/Am473ur/HexQBot
HexQBot:
  enabled: false
  api: http://xxx.xxxxxx.com/send
  qqgroup: 000000000
  timeout: 2

# Server酱
# https://sct.ftqq.com/
ServerChan:
  enabled: false
  sendkey: xxxxxxxxxxxxxxxxxxxxxx
  timeout: 2


############### SiteEnable ###############

# 棱角社区
# https://forum.ywhack.com/forum-59-1.html
EdgeForum:
  enabled: true

# 先知安全技术社区
# https://xz.aliyun.com/
XianZhi:
  enabled: true

# SeebugPaper（知道创宇404实验室）
# https://paper.seebug.org/
SeebugPaper:
  enabled: true

# 安全客
# https://www.anquanke.com/
Anquanke:
  enabled: true

```


## Contributing


如果您有高质量的安全社区网站希望被爬取，或者想推荐被广泛使用的推送机器人，欢迎联系我`leonsec[at]h4ck.fun`或提交[issue](https://github.com/Le0nsec/SecCrawler/issues)和[PR](https://github.com/Le0nsec/SecCrawler/pulls)。



## License


[GNU General Public License v3.0](https://github.com/Le0nsec/SecCrawler/blob/main/LICENSE)