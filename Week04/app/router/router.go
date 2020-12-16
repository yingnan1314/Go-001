package router

import (
	"github.com/gin-gonic/gin"
	"ppwords/app/service/words"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/get/words", words.GetWordsByUser)
	r.GET("/update/words", words.UpdateWordsByUser)

	return r
}