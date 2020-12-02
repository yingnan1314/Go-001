package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qyn/dao"
	"qyn/models"
	"qyn/models/request"
	"qyn/models/response"
	"qyn/util/validator"
)

type studyService struct {
	dao dao.StudyDao
}

type StudyService interface {
	GetListById(page, pageSize int) ([]models.Study,error)
}

func (s *studyService) GetListById(ctx *gin.Context) {
	req := &request.PageInfo{}
	if err := ctx.BindJSON(&req); err != nil {
		panic(err)
	}
	fmt.Println(req)
	err := validator.Validate(req)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%s", err), ctx)
		return
	}
	res,err := s.dao.GetListById(req.Page, req.PageSize)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%s", err), ctx)
	}

	response.OkWithData(res,ctx)

}
