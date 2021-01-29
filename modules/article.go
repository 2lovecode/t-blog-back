package modules

import (
	"go.mongodb.org/mongo-driver/bson"
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


type ArticleListReq struct {
	Page int64		`form:"pageNum" json:"page"`
	PageSize int64	`form:"pageSize" json:"pageSize"`
	TagID string	`form:"tagID" json:"tagID"`
}

type ArticleListResp struct {
	Page utils.PaginationData `json:"pagination"`
	TagID string `json:"tagID,omitempty"`
	TagName string `json:"tagName,omitempty"`
	List []models.Article `json:"list"`
}

// GetArticleList 获取列表
func GetArticleList(c *gin.Context) {
	articleListReq := ArticleListReq{}

	err := c.ShouldBindQuery(&articleListReq)

	if err == nil {
		article := &models.Article{}
		tagName := ""

		filter := bson.D{}

		if articleListReq.TagID != "" {
			filter = append(filter, bson.E{
				Key: "tagID",
				Value: articleListReq.TagID,
			})
			tag := &models.Tag{}
			t, _ := tag.FindByID(c, articleListReq.TagID)
			tagName = t.Name
		}

		pagination := utils.NewPagination(10)
		pagination.SetPage(articleListReq.Page)
		pagination.SetSize(articleListReq.PageSize)

		articleList := article.GetArticles(c, pagination, filter)

		pageData := pagination.GetPaginationData()

		utils.SuccessJSON(c, ArticleListResp{
			Page:     pageData,
			TagID:    articleListReq.TagID,
			TagName:  tagName,
			List:     articleList,
		})
	} else {
		utils.FailureJSON(c, err)
	}

}

// GetArticleDetail 获取文章详情
func GetArticleDetail(c *gin.Context) {
	code := e.Success

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
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
