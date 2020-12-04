package middleware

import (
	"errors"
	"t-blog-back/models"
	"t-blog-back/pkg/setting"
	"t-blog-back/pkg/utils"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const identityKey = "AuthCode"
const realm = "public"

// LoginForm 登录验证
type LoginForm struct {
	UserName string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
}

// AuthData user data
type AuthData struct {
	AuthorCode string `json:"authorCode"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
}

// NewAuthMiddleware 登录中间件
func NewAuthMiddleware() (auth *jwt.GinJWTMiddleware, err error) {
	auth, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       realm,
		Key:         []byte(setting.AuthCfg.SecretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*AuthData); ok {
				return jwt.MapClaims{
					identityKey: v.AuthorCode,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &AuthData{
				AuthorCode: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals LoginForm
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userName := loginVals.UserName
			passWord := loginVals.PassWord
			user := models.User{}
			if user.FindUserByName(c, userName); err == nil && bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(passWord)) == nil {
				// 成功后，生成一个有过期时间的授权码
				authorCode := utils.GenUniqueID()
				l := &models.Login{}
				login, _ := l.FindLoginByAuthorID(c, user.AuthorID)
				if login == nil {
					login = &models.Login{
						AuthorID:  user.AuthorID,
						AuthCode:  authorCode,
						Expire:    3600,
						StartTime: time.Now(),
					}
					login.AddLogin(c)
				} else {
					login.UpdateLoginAuthCode(c, authorCode)
				}
				return &AuthData{
					AuthorCode: authorCode,
					Name:       user.Name,
					Avatar:     user.Avatar,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*AuthData); ok {
				if v.AuthorCode != "" {
					l := &models.Login{}
					login, err := l.FindUserByAuthCode(c, v.AuthorCode)
					if err == nil && login != nil {
						return true
					}
				}
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			utils.FailureJSONWithHTTPCode(c, errors.New(message), code)
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			utils.SuccessJSON(c, map[string]interface{}{
				"token":  token,
				"expire": expire,
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			utils.SuccessJSON(c, map[string]interface{}{
				"token":  token,
				"expire": expire,
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			utils.SuccessJSON(c, nil)
		},
		TokenLookup:   "header: TankBlog-Token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err == nil && auth != nil {
		err = auth.MiddlewareInit()
	}
	return
}
