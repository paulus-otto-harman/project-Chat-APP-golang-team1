package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"project/api-gateway/infra"
)

func main() {
	var err error

	var ctx *infra.ServiceContext
	if ctx, err = infra.NewServiceContext(); err != nil {
		log.Fatal("can't init service context %w", err)
	}

	router := gin.Default()
	router.GET("test", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
	log.Println(fmt.Sprintf("starting server on %s:%s", ctx.Cfg.ServerIp, ctx.Cfg.ServerPort))
	if err = router.Run(fmt.Sprintf("%s:%s", ctx.Cfg.ServerIp, ctx.Cfg.ServerPort)); err != nil {
		return
	}
}
