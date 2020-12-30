package routers

import jwt "github.com/appleboy/gin-jwt/v2"

type AppJWTMiddleware struct {
	*jwt.GinJWTMiddleware
}
