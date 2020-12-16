package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ppwords/backend/service"
	"ppwords/dao"
	"ppwords/models/request"
	"ppwords/models/response"
	"ppwords/util/validator"
)

type WordsController struct {
	wordsService service.WordsService
}


func (w *WordsController) GetWordsList(c *gin.Context) {
	req := &request.PageInfo{}
	if err := c.BindJSON(&req); err != nil {
		panic(err)
	}
	fmt.Println(req)
	err := validator.Validate(req)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%s", err), c)
		return
	}

	res,err := w.wordsService.GetWordsList(req.Page, req.PageSize)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%s", err), c)
	}

	response.OkWithData(res,c)
}

func (w *WordsController) CreateWord(c *gin.Context) {
	word := dao.Words{}
	if err := c.BindJSON(&word); err != nil {
		response.FailWithMessage("Field  is Empty!", c)
		return
	}
	if word.Grade == 0 {
		response.FailWithMessage("Field Grade is Unique!", c)
		return
	}
	if err := w.wordsService.CreateWord(word); err == nil {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func (w *WordsController) UpdateWord(c *gin.Context) {
	word := dao.Words{}
	if err := c.BindJSON(&word); err != nil {
		response.FailWithMessage("Field  is Empty!", c)
		return
	}
	if word.Grade == 0 {
		response.FailWithMessage("Field Grade is Unique!", c)
		return
	}
	fmt.Println(word)
	if err := w.wordsService.UpdateWord(word); err == nil {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}