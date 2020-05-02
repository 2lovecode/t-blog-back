package routers

import (
	"github.com/gin-gonic/gin"
	"t-blog-back/middleware"
	"t-blog-back/pkg/setting"
	v1 "t-blog-back/routers/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	apiV1 := r.Group("/api/v1")
	{
		//模块
		apiV1.GET("/modules", v1.GetModuleList)
		apiV1.GET("/module/:id", v1.GetModuleDetail)
		apiV1.POST("/module", v1.AddModule)
		apiV1.PUT("/module/:id", v1.EditModule)
		apiV1.DELETE("/module/:id", v1.SoftDeleteModule)

		//文章
		apiV1.GET("/articles", v1.GetArticleList)
		apiV1.GET("/article/:id", v1.GetArticleDetail)
		apiV1.POST("/article", v1.AddArticle)
		apiV1.PUT("/article/:id", v1.EditArticle)
		apiV1.DELETE("/article/:id", v1.SoftDeleteArticle)

		//标签
		apiV1.GET("/tags", v1.GetTagList)
		apiV1.GET("/tag/:id", v1.GetTagDetail)
		apiV1.POST("/tag", v1.AddTag)
		apiV1.PUT("/tag/:id", v1.EditTag)
		apiV1.DELETE("/tag/:id", v1.SoftDeleteTag)

	}

	return r
}