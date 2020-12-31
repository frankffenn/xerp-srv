package routers

import (
	"github.com/frankffenn/xerp-srv/errors"
	"github.com/gin-gonic/gin"
)

func ResponseSuccess(data map[string]interface{}) gin.H {
	return gin.H{
		"success": true,
		"data":    data,
	}
}

func ResponseFailWithError(err *errors.Error) gin.H {
	return gin.H{
		"success": false,
		"code":    err.Code,
		"msg":     err.Message,
	}
}
