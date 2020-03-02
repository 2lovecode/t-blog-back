package v1

import (
	"github.com/gin-gonic/gin"
)

func GetArticleList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "article list",
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
