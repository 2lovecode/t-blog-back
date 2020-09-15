package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/e"
	"t-blog-back/pkg/utils"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TankBlog-Token")
		if token == "" {
			utils.Abort(c, http.StatusUnauthorized, e.ErrorInvalidToken, nil)
			return
		}

		user := models.User{}
		u, err := user.FindUserByToken(token)

		if err != nil || u.Name == "" {
			utils.Abort(c, http.StatusUnauthorized, e.ErrorInvalidToken, nil)
			return
		}

		jsonP := utils.GetJsonParser()
		uInfo, err := jsonP.Marshal(&u)

		if err != nil {
			utils.Abort(c, http.StatusUnauthorized, e.ErrorInvalidToken, nil)
			return
		}
		c.Set("user-info", string(uInfo))
		c.Next()
	}
}
