package modules

import (
	"t-blog-back/models"
	"t-blog-back/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// AddCategoryReq 参数验证
type AddCategoryReq struct {
	Name string `form:"name" json:"name" binding:"required"`
}

// AddCategoryResp 返回
type AddCategoryResp struct {
	CategoryID string `json:"categoryID"`
}

// GetCategoryList 分类列表
func GetCategoryList(c *gin.Context) {

}

// GetCategoryDetail 分类详情
func GetCategoryDetail(c *gin.Context) {

}

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var req AddCategoryReq
	err := c.ShouldBindJSON(&req)

	if err == nil {
		category := &models.Category{}

		if err = category.FindByName(c, req.Name); err == nil {
			resp := AddCategoryResp{}
			if category.IsEmpty() {
				category.Name = req.Name
				category.ID = utils.GenUniqueID()
				category.State = models.CategoryStateNormal
				category.AddTime = time.Now()
				category.ModifyTime = time.Now()
				_, err = category.AddCategory(c)
				if err == nil {
					resp.CategoryID = category.ID
				}
			} else {
				resp.CategoryID = category.ID
			}
			if err == nil && resp.CategoryID != "" {
				utils.SuccessJSON(c, resp)
				return
			}
		}
	}
	utils.FailureJSON(c, err)
	return
}

// EditCategory 修改分类
func EditCategory(c *gin.Context) {

}

// SoftDeleteCategory 删除分类
func SoftDeleteCategory(c *gin.Context) {

}
