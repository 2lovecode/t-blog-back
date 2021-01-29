package modules

import (
	"go.mongodb.org/mongo-driver/bson"
	"t-blog-back/models"
	"t-blog-back/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// AddTagReq 请求
type AddTagReq struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type TagListReq struct {
	Page int64		`form:"pageNum" json:"page"`
	PageSize int64	`form:"pageSize" json:"pageSize"`
}

type TagListResp struct {
	Page utils.PaginationData `json:"pagination"`
	List []models.Tag `json:"list"`
}

// GetTagList 获取标签列表
func GetTagList(c *gin.Context) {
	tag := &models.Tag{}
	tagReq := TagListReq{}

	err := c.ShouldBindQuery(&tagReq)

	if err == nil {
		pagination := utils.NewPagination(10)
		pagination.SetPage(tagReq.Page)
		pagination.SetSize(tagReq.PageSize)
		data, err := tag.GetTags(c, pagination, bson.D{})

		pageData := pagination.GetPaginationData()

		if err == nil {
			utils.SuccessJSON(c, TagListResp{
				Page: pageData,
				List: data,
			})
			return
		}
	}

	utils.FailureJSON(c, err)
}

// GetTagDetail 获取详情
func GetTagDetail(c *gin.Context) {

}

// AddTag 添加
func AddTag(c *gin.Context) {
	var req AddTagReq
	err := c.ShouldBindJSON(&req)

	data := make(map[string]string)

	if err == nil {
		tag := &models.Tag{}

		if _, err0 := tag.FindByName(c, req.Name); err0 == nil {
			utils.SuccessJSON(c, data)
		} else {
			tag.Name = req.Name
			tag.ID = utils.GenUniqueID()
			tag.State = models.TagStateNormal
			tag.AddTime = time.Now()
			tag.ModifyTime = time.Now()
			if _, err = tag.AddTag(c); err == nil {
				data = map[string]string{
					"id": tag.ID,
				}
				utils.SuccessJSON(c, data)
			}
		}

	} else {
		utils.FailureJSON(c, err)
	}
}

// EditTag 修改
func EditTag(c *gin.Context) {

}

// SoftDeleteTag 删除
func SoftDeleteTag(c *gin.Context) {

}
