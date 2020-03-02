package routers

import (
	"github.com/gin-gonic/gin"
	"t-blog-back/pkg/setting"
	v1 "t-blog-back/routers/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	apiV1 := r.Group("/api/v1")
	{
		//分类
		apiV1.GET("/categories", v1.GetCategoryList)
		apiV1.GET("/category/:id", v1.GetCategoryDetail)
		apiV1.POST("/category", v1.AddCategory)
		apiV1.PUT("/category/:id", v1.EditCategory)
		apiV1.DELETE("/category/:id", v1.SoftDeleteCategory)

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