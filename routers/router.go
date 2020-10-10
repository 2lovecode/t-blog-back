package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"t-blog-back/middleware"
	"t-blog-back/modules"
	"t-blog-back/pkg/e"
	"t-blog-back/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		code := e.Success

		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"data" : "test message",
		})
	})

	r.POST("/login", modules.Login)

	apiV1 := r.Group("/api/v1", middleware.Login())
	{
		//分类
		apiV1.GET("/categories", modules.GetCategoryList)
		apiV1.GET("/category/:id", modules.GetCategoryDetail)
		apiV1.POST("/category", modules.AddCategory)
		apiV1.PUT("/category/:id", modules.EditCategory)
		apiV1.DELETE("/category/:id", modules.SoftDeleteCategory)

		//标签
		apiV1.GET("/tags", modules.GetTagList)
		apiV1.GET("/tag/:id", modules.GetTagDetail)
		apiV1.POST("/tag", modules.AddTag)
		apiV1.PUT("/tag/:id", modules.EditTag)
		apiV1.DELETE("/tag/:id", modules.SoftDeleteTag)

		//文章
		apiV1.GET("/articles", modules.GetArticleList)
		apiV1.GET("/article/:id", modules.GetArticleDetail)
		apiV1.POST("/article", modules.AddArticle)
		apiV1.PUT("/article/:id", modules.EditArticle)
		apiV1.DELETE("/article/:id", modules.SoftDeleteArticle)

	}

	return r
}