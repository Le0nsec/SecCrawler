package api

import (
	"SecCrawler/api/controllers"
	"SecCrawler/config"
	"SecCrawler/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	setCors(r)
	api := r.Group("/api", auth)

	public := api.Group("/crawler")
	{
		public.GET("/getArticles/:site", controllers.GetArticles)
	}
}

func setCors(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	r.Use(cors.New(config))
}

func auth(c *gin.Context) {
	key := c.GetHeader("Authorization")

	if key != config.Cfg.Api.AuthKey {
		utils.ErrorStrResp(c, utils.INVALID_AUTH_KEY, "Invalid auth key")
		return
	}
	c.Next()
}
