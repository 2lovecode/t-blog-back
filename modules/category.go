package modules

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/e"
	"t-blog-back/pkg/utils"
	"time"
)

type AddCategoryReq struct {
	Name string `form:"name" binding:"required"`
}

type AddCategoryResp struct {

}

//分类列表
func GetCategoryList(c *gin.Context) {

}

//分类详情
func GetCategoryDetail(c *gin.Context) {

}

//添加分类
func AddCategory(c *gin.Context) {
	var req AddCategoryReq
	err := c.ShouldBindJSON(&req)

	var code e.RCode
	code = e.ErrorInvalidParams
	eMsg := ""
	data := make(map[string]string)

	if err == nil {
		category := &models.Category{}

		category.Name = req.Name
		category.State = models.CategoryStateNormal
		category.AddTime = time.Now()
		category.ModifyTime = time.Now()

		result, err := category.AddCategory()

		if err == nil {
			data = map[string]string{
				"id": result.InsertedID.(string),
			}
		} else {
			eMsg = err.Error()
		}
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"error": eMsg,
			"data" : data,
		})
	} else {
		utils.AbortWithMessage(c, http.StatusBadRequest, e.Error, err.Error(), nil)
	}
}

//修改分类
func EditCategory(c *gin.Context) {

}

//删除分类
func SoftDeleteCategory(c *gin.Context) {

}