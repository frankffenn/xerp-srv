package main

import (
	"net/http"
	"time"

	"github.com/frankffenn/go-utils/log"
	"github.com/frankffenn/xerp-srv/config"
	"github.com/frankffenn/xerp-srv/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal("init config: %v", err)
	}

	gin.SetMode(config.App.Mode)

	router := gin.New()

	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	c.AllowHeaders = []string{"*"}
	c.AllowMethods = []string{"GET", "POST", "PUT", "OPTION"}
	c.MaxAge = time.Hour
	router.Use(cors.New(c))

	if err := routers.Init(router); err != nil {
		log.Fatal("init routers: %v", err)
	}

	srv := &http.Server{
		Addr:    config.App.ListenAddress,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen err %v", err)
	}
}
