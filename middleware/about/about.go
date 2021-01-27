package about

import (
	"t-blog-back/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AboutForm struct {
	CreateTime string `form:"create_time" binding:"required,datetime=2006-01-02"`
}

//表单验证
func AboutValidate() gin.HandlerFunc {
	return func(c *gin.Context) {

		var form AboutForm

		err := c.ShouldBindQuery(&form)

		if err != nil {
			utils.FailureJSON(c, err)
			return
		}
		c.Next()
	}
}
