package initRouter

import (
	"GinHello/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.Static("/statics", "./statics/")
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	// 添加 user
	apiRouter := router.Group("/v1")
	{
		apiRouter.GET("/getadress", handler.Getadress)
	}
	return router
}
