package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/frankffenn/go-utils/log"
	"github.com/frankffenn/xerp-srv/errors"
	"github.com/frankffenn/xerp-srv/services/users"
	"github.com/frankffenn/xerp-srv/services/users/db"
	usrmod "github.com/frankffenn/xerp-srv/services/users/mod"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type AppJWTMiddleware struct {
	*jwt.GinJWTMiddleware
}

func JwtAuthenticator(c *gin.Context) (interface{}, error) {
	var login usrmod.LoginVar
	if err := c.ShouldBind(&login); err != nil {
		return "", errors.ErrInvalidRequestParams
	}

	switch login.LoginType {
	case usrmod.GuestLogin:
		return guestAuth(c, login.Username)
	case usrmod.PhoneLogin:
		return phoneAuth(c, login.Username, login.Password, false)
	case usrmod.WechatLogin:
		log.Info("implement me")
	default:
		return "", errors.ErrUnsupportedLoginType
	}

	return nil, nil
}

func guestAuth(ctx context.Context, username string) (interface{}, error) {
	found, err := users.GetUserInfoFromUsername(ctx, username)
	if err != nil && !xerrors.Is(err, db.ErrUserNotFound) {
		return nil, err
	}

	if found != nil {
		return &usrmod.LoginResp{GUID: found.GUID, UserID: found.ID}, nil
	}

	u := &usrmod.User{
		Username:  username,
		LoginType: usrmod.GuestLogin,
		GUID:      uuid.New().String(),
	}
	if err := users.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &usrmod.LoginResp{GUID: u.GUID, UserID: u.ID}, nil
}

func phoneAuth(ctx context.Context, username, password string, checkAdmin bool) (interface{}, error) {
	return nil, nil
}

func wechatAuth(username string) (interface{}, error) {
	return nil, nil
}

func JwtPayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*usrmod.LoginResp); ok {
		return jwt.MapClaims{
			"guid":    v.GUID,
			"user_id": v.UserID,
			"level":   v.Level,
		}
	}
	return jwt.MapClaims{}
}

func JwtIdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &usrmod.LoginResp{
		GUID:   claims["guid"].(string),
		UserID: uint64(claims["user_id"].(float64)),
		Level:  uint64(claims["level"].(float64)),
	}
}

func JwtLoginResponse(c *gin.Context, code int, token string, expire time.Time) {
	jToken, err := AuthUserMiddleware.ParseTokenString(token)
	claims := jwt.ExtractClaimsFromToken(jToken)
	userId := uint64(claims["user_id"].(float64))

	auth := &struct {
		CurrToken string `json:"curr_token"`
		LastToken string `json:"last_token"`
	}{
		CurrToken: token,
	}

	_, err = json.Marshal(auth)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithError(errors.ErrTokenCreateFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(map[string]interface{}{
		"user_id":   userId,
		"token":     token,
		"expire":    expire.Format(time.RFC3339),
		"expire_ts": expire.Unix(),
	}))
}
