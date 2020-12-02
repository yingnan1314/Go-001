package dao

import (
	"database/sql"
	"errors"
	"log"
	"qyn/models"
	xerrors "github.com/pkg/errors"
)

type StudyDao struct {
	*sql.DB
}

func (s *StudyDao) GetListById(page, pageSize int) ([]models.Study,error) {
	if page < 1 {
		page = 1
	}
	list := make([]models.Study, 0)
	study := new(models.Study)
	res,err := DB.Query("SELECT `id`, `title`, `description` FROM `study` LIMIT ? OFFSET ?", pageSize, pageSize * (page - 1))
	if err != nil {
		log.Fatal(xerrors.Wrap(err, "数据查询失败"))
		return nil, errors.New("没有查询到数据")
	}
	for res.Next() {
		if err := res.Scan(&study.Id, &study.Title, &study.Description); err == nil {
			list = append(list, *study)
		}
	}
	if list == nil {
		return nil, errors.New("没有查询到数据")
	}
	return list,nil
}
