package modules

import (
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/e"
	"t-blog-back/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// AddTagReq 请求
type AddTagReq struct {
	Name string `form:"name" json:"name" binding:"required"`
}

// GetTagList 获取标签列表
func GetTagList(c *gin.Context) {
	maps := make(map[string]interface{})
	code := e.Success

	tag := &models.Tag{}

	data := tag.GetTags(c, 1, 10, maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
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

		if _, err0 := tag.FindByName(req.Name); err0 == nil {
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
