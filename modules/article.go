package modules

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/e"
)

func GetArticleList(c *gin.Context) {
	articleList := models.GetArticleList()

	code := e.Success

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : articleList,
	})
}

//获取文章详情
func GetArticleDetail(c *gin.Context) {

}

//添加文章
func AddArticle(c *gin.Context) {

}

//修改文章
func EditArticle(c *gin.Context) {

}

//软删除文章
func SoftDeleteArticle(c *gin.Context) {

}
