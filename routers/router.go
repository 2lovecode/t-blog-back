package routers

import (
	"errors"
	"fmt"
	"net/http"
	"t-blog-back/middleware"
	"t-blog-back/modules"
	"t-blog-back/pkg/e"
	"t-blog-back/pkg/setting"
	"t-blog-back/pkg/utils"

	jwt "github.com/appleboy/gin-jwt/v2"
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

	auth, _ := middleware.NewAuthMiddleware()

	r.POST("/api/login", auth.LoginHandler)
	r.GET("/api/refresh_token", auth.RefreshHandler)

	r.NoRoute(auth.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		msg := fmt.Sprintf("NoRoute claims: %#v\n", claims)
		utils.FailureJSONWithHTTPCode(c, errors.New(msg), http.StatusNotFound)
	})

	apiFrontendV1 := r.Group("/api/v1", middleware.Cors())
	{
		// app详情
		apiFrontendV1.GET("/about", modules.About)
		// 文章列表
		apiFrontendV1.GET("/article-list", modules.GetArticleList)
		// 文章详情
		apiFrontendV1.GET("/article-detail", modules.GetArticleDetail)
		// 分类列表
		apiFrontendV1.GET("/category-list", modules.GetCategoryList)
	}

	apiBackendV1 := r.Group("/api/backend/v1", auth.MiddlewareFunc())
	{
		//分类
		apiBackendV1.GET("/categories", modules.GetCategoryList)
		apiBackendV1.GET("/category/:id", modules.GetCategoryDetail)
		apiBackendV1.POST("/add-category", modules.AddCategory)
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
