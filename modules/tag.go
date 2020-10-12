package modules

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/e"
	"t-blog-back/pkg/utils"
	"time"
)

type AddTagReq struct {
	Name string `form:"name" json:"name" binding:"required"`
}

func GetTagList(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.Success

	data["lists"] = models.GetTags(1, 2, maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func GetTagDetail(c *gin.Context) {

}

func AddTag(c *gin.Context) {
	var req AddTagReq
	err := c.ShouldBindJSON(&req)

	var code e.RCode
	code = e.ErrorInvalidParams
	eMsg := ""
	data := make(map[string]string)

	if err == nil {
		tag := &models.Tag{}

		if _, err0 := tag.FindByName(req.Name); err0 == nil {
			utils.Success(c, code, eMsg, data)
		} else {
			tag.Name = req.Name
			tag.ID = utils.GenUniqueID()
			tag.State = models.TagStateNormal
			tag.AddTime = time.Now()
			tag.ModifyTime = time.Now()
			if _, err = tag.AddTag(); err == nil {
				data = map[string]string {
					"id": tag.ID,
				}
				utils.Success(c, code, eMsg, data)
			}
		}

	} else {
		utils.AbortWithMessage(c, http.StatusBadRequest, e.Error, err.Error(), nil)
	}
}

func EditTag(c *gin.Context) {

}

func SoftDeleteTag(c *gin.Context) {

}
