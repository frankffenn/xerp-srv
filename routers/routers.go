package routers

import (
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/frankffenn/xerp-srv/config"
	"github.com/gin-gonic/gin"
)

var (
	TokenExpired        = time.Hour * 24
	TokenRefreshTimeout = time.Hour * 24 * 30
	AuthUserMiddleware  *AppJWTMiddleware
)

func Init(router *gin.Engine) error {
	jwtAuthUserMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "User",
		Key:         []byte(config.App.JWTScrect),
		Timeout:     TokenExpired,
		MaxRefresh:  TokenRefreshTimeout,
		IdentityKey: "guid",
	})
	if err != nil {
		return err
	}

	AuthUserMiddleware = &AppJWTMiddleware{
		GinJWTMiddleware: jwtAuthUserMiddleware,
	}

	apiV1 := router.Group("/api/v1")
	apiV1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{"pong": true})
	})

	usr := apiV1.Group("/user")
	usr.POST("/login", AuthUserMiddleware.LoginHandler)
	router.Use(AuthUserMiddleware.MiddlewareFunc())
	return nil
}
