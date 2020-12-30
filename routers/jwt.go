package routers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/frankffenn/xerp-srv/errors"
	"github.com/frankffenn/xerp-srv/go-utils/log"
	user "github.com/frankffenn/xerp-srv/services/users/mod"
	"github.com/gin-gonic/gin"
)

type AppJWTMiddleware struct {
	*jwt.GinJWTMiddleware
}

func JwtAuthenticator(c *gin.Context) (interface{}, error) {
	var login user.LoginVar
	if err := c.ShouldBind(&login); err != nil {
		return "", errors.InvalidRequestParams
	}

	switch login.LoginType {
	case user.GuestLogin:
		return guestAuth()
	case user.PhoneLogin:
		return phoneAuth(login.Username, login.Password, false)
	case user.WechatLogin:
		log.Info("implement me")
	default:
		return "", errors.ErrUnsupportedLoginType
	}
}

func guestAuth() (interface{}, error) {
	return &user.loginResp{}, nil
}

func phoneAuth(username, password string, checkAdmin bool) (interface{}, error) {
	return nil, nil
}

func wechatAuth(username string) (interface{}, error) {
	return nil, nil
}
