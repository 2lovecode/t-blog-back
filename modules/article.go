package modules

import (
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/e"
	"t-blog-back/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AddArticleReq struct {
	AuthorID     string    `form:"authorID" json:"authorID" binding:"required"`
	Title        string    `form:"title" json:"title" binding:"required"`
	Image        string    `form:"image" json:"image" binding:"required"`
	Summary      string    `form:"summary" json:"summary" binding:"required"`
	Content		 string    `form:"content" json:"content" binding:"required"`
	Tags         []string  `form:"tags" json:"tags" binding:"required"`
}

// GetArticleList 获取列表
func GetArticleList(c *gin.Context) {
	article := &models.Article{}

	articleList := article.GetArticles(c, 1, 10, nil)

	code := e.Success

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": articleList,
	})
}

// GetArticleDetail 获取文章详情
func GetArticleDetail(c *gin.Context) {
	code := e.Success

	user, _ := utils.GetLoginUserInfo(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": user,
	})
}

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	addReq := AddArticleReq{}

	err := c.ShouldBindJSON(&addReq)

	if err == nil {
		article := &models.Article{}
		article.ArticleID = utils.GenUniqueID()
		article.Title = addReq.Title
		article.Image = addReq.Image
		article.AuthorID = addReq.AuthorID
		article.Content = addReq.Content
		article.Summary = addReq.Summary
		article.Tags = addReq.Tags

		_, err = article.AddArticle(c)
		if err == nil {
			utils.SuccessJSON(c, nil)
			return
		}
	}
	utils.FailureJSON(c, err)
}

// EditArticle 修改文章
func EditArticle(c *gin.Context) {

}

// SoftDeleteArticle 软删除文章
func SoftDeleteArticle(c *gin.Context) {

}
