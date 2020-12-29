package main

import (
	"net/http"

	"github.com/frankffenn/xerp-srv/config"
	"github.com/frankffenn/xerp-srv/go-utils/log"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal("init config:", err)
	}

	gin.SetMode(config.App.Mode)

	router := gin.New()
	apiV1 := router.Group("/api/v1")
	apiV1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{"pong": true})
	})

	srv := &http.Server{
		Addr:    config.App.ListenAddress,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("server err", err)
	}
}
