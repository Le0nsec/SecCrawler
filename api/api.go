package api

import (
	"SecCrawler/api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	setCors(r)
	api := r.Group("/api")

	public := api.Group("/crawler")
	{
		public.GET("/getArticles/:site", controllers.GetArticles)
	}
}

func setCors(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization", "range")
	r.Use(cors.New(config))
}
