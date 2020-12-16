package router

import (
	"github.com/gin-gonic/gin"
	controller2 "ppwords/backend/controller"
)

func InitRouter() *gin.Engine {
	var controller controller2.WordsController
	r := gin.Default()
	r.POST("/words/getList", controller.GetWordsList)
	r.POST("/words/create", controller.CreateWord)
	r.POST("/words/update", controller.UpdateWord)

	return r
}