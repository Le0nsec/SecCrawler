package controllers

import (
	"SecCrawler/register"
	"SecCrawler/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	siteName := c.Params.ByName("site")
	crawler, ok := register.GetCrawler(siteName)
	if !ok {
		utils.ErrorStrResp(c, utils.SITE_NOT_FOUND, "The site is not open or does not exist")
		return
	}
	fmt.Printf("[*] api call [%s]\n", crawler.Config().Name)
	result, err := crawler.Get()
	if err != nil {
		utils.ErrorResp(c, utils.ARTICLE_NOT_FOUND, err)
		return
	}
	utils.SuccessResp(c, result)
}
