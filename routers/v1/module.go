package v1

import (
	"database/sql"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/e"
	"t-blog-back/pkg/utils"
)

type AddModuleReq struct {
	Name string `json:"name"`
}
//获取模块列表
func GetModuleList(c *gin.Context) {
	sql.Open("sqlite3", "")
}

func GetModuleDetail(c *gin.Context) {

}

//添加模块
func AddModule(c *gin.Context) {
	var req AddModuleReq
	err := c.ShouldBindJSON(&req)

	var code e.RCode
	code = e.ErrorInvalidParams
	eMsg := ""
	data := make(map[string]string)

	if err == nil {
		name := req.Name
		valid := validation.Validation{}

		valid.Required(name, "name").Message("名字不能为空")
		valid.MaxSize(name, 100, "name").Message("名字最长为100字符")

		if !valid.HasErrors() {
			if !models.ModuleExistByName(name) {
				code = e.Success
				models.AddModule(name)
			} else {
				code = e.ErrorExistModule
			}
		} else {
			eMsg = utils.GetFirstErrorMessage(valid)
		}
	} else {
		eMsg = "参数必须为json格式"
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"error": eMsg,
		"data" : data,
	})

}

//修改模块
func EditModule(c *gin.Context) {

}

//软删除类
func SoftDeleteModule(c *gin.Context) {

}
