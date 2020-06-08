package modules

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/e"
)

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

}

func EditTag(c *gin.Context) {

}

func SoftDeleteTag(c *gin.Context) {

}
