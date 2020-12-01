package routers

import (
	"net/http"
	"t-blog-back/middleware"
	"t-blog-back/modules"
	"t-blog-back/pkg/e"
	"t-blog-back/pkg/setting"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	gin.SetMode(setting.RunMode)

	r.GET("/api/test", func(c *gin.Context) {
		code := e.Success

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": "test message",
		})
	})

	r.POST("/api/no-login-token", modules.NoLoginToken)
	r.POST("/api/login", modules.Login)
	r.POST("/api/refresh-token", modules.RefreshToken)

	apiFrontendV1 := r.Group("/api/v1", middleware.Cors(), middleware.Login())
	{
		apiFrontendV1.POST("no-login-token")
	}

	apiBackendV1 := r.Group("/api/backend/v1", middleware.Login())
	{
		//分类
		apiBackendV1.GET("/categories", modules.GetCategoryList)
		apiBackendV1.GET("/category/:id", modules.GetCategoryDetail)
		apiBackendV1.POST("/category", modules.AddCategory)
		apiBackendV1.PUT("/category/:id", modules.EditCategory)
		apiBackendV1.DELETE("/category/:id", modules.SoftDeleteCategory)

		//标签
		apiBackendV1.GET("/tags", modules.GetTagList)
		apiBackendV1.GET("/tag/:id", modules.GetTagDetail)
		apiBackendV1.POST("/tag", modules.AddTag)
		apiBackendV1.PUT("/tag/:id", modules.EditTag)
		apiBackendV1.DELETE("/tag/:id", modules.SoftDeleteTag)

		//文章
		apiBackendV1.GET("/articles", modules.GetArticleList)
		apiBackendV1.GET("/article/:id", modules.GetArticleDetail)
		apiBackendV1.POST("/article", modules.AddArticle)
		apiBackendV1.PUT("/article/:id", modules.EditArticle)
		apiBackendV1.DELETE("/article/:id", modules.SoftDeleteArticle)

	}

	return r
}
