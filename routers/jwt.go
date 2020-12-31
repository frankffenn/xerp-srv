package routers

import (
	"context"
	"fmt"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/frankffenn/xerp-srv/errors"
	"github.com/frankffenn/xerp-srv/go-utils/log"
	"github.com/frankffenn/xerp-srv/services/users"
	user "github.com/frankffenn/xerp-srv/services/users/mod"
	usrmod "github.com/frankffenn/xerp-srv/services/users/mod"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AppJWTMiddleware struct {
	*jwt.GinJWTMiddleware
}

func JwtAuthenticator(c *gin.Context) (interface{}, error) {
	var login user.LoginVar
	if err := c.ShouldBind(&login); err != nil {
		return "", errors.ErrInvalidRequestParams
	}

	switch login.LoginType {
	case user.GuestLogin:
		return guestAuth(c, login.Username)
	case user.PhoneLogin:
		return phoneAuth(c, login.Username, login.Password, false)
	case user.WechatLogin:
		log.Info("implement me")
	default:
		return "", errors.ErrUnsupportedLoginType
	}

	return nil, nil
}

func guestAuth(ctx context.Context, username string) (interface{}, error) {
	u := &usrmod.User{
		Username:  fmt.Sprintln(username),
		LoginType: usrmod.GuestLogin,
		GUID:      uuid.New().String(),
	}
	if err := users.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &user.LoginResp{GUID: u.GUID, UserID: u.ID}, nil
}

func phoneAuth(ctx context.Context, username, password string, checkAdmin bool) (interface{}, error) {
	return nil, nil
}

func wechatAuth(username string) (interface{}, error) {
	return nil, nil
}
