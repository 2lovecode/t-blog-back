package middleware

import (
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/e"
	"t-blog-back/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Login 登录验证
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TankBlog-Token")
		if token == "" {
			utils.Abort(c, http.StatusUnauthorized, e.FailureInvalidToken, nil)
			return
		}

		user := models.User{}
		u, err := user.FindUserByToken(token)

		if err != nil || u.Name == "" {
			utils.Abort(c, http.StatusUnauthorized, e.FailureInvalidToken, nil)
			return
		}

		jsonP := utils.GetJsonParser()
		uInfo, err := jsonP.Marshal(&u)

		if err != nil {
			utils.Abort(c, http.StatusUnauthorized, e.FailureInvalidToken, nil)
			return
		}
		c.Set("user-info", string(uInfo))
		c.Next()
	}
}
