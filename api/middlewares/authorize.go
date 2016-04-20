package middlewares

import (
	"net/http"
	"strings"

	"github.com/iguiyu/demo/api/configs"
	"github.com/iguiyu/demo/misc/global"
	"github.com/iguiyu/demo/misc/helpers"
	"github.com/iguiyu/demo/misc/oauth"

	"github.com/kobeld/goutils"

	"github.com/gin-gonic/gin"

	userReqres "github.com/iguiyu/microservices/user/reqres"
	"gopkg.in/mgo.v2/bson"
)

func OAuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Check if the Authorization header exists
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		// check if the Bearer Token exists
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
			c.Next()
			return
		}

		accessData, err := oauth.GetAccess(parts[1])
		if err != nil {
			if err != global.ErrInvalidToken { // Should be a system error
				goutils.PrintStackAndError(err)
			}
			c.Next()
			return
		}

		// Check the Access data
		if accessData.IsExpired() || accessData.UserData == nil {
			c.Next()
			return
		}

		var (
			userId = accessData.UserData.(string)
			req    = &userReqres.UserReq{Id: bson.ObjectIdHex(userId)}
			res    = &userReqres.UserRes{}
		)

		if helpers.CallRpcServiceWithContext(c, configs.AppConf, userReqres.SERVICE_NAME, userReqres.RPC_USER_SHOW, req, res) {
			return
		}

		c.Set("me", res.User)

		c.Next()
	}
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, exists := c.Get("me")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
