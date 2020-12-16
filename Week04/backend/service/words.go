package service

import (
	"errors"
	"ppwords/dao"
)

type WordsService struct {
	dao dao.Words
}

// GetWordsList get words list by page
func (w *WordsService) GetWordsList(page, pageSize int) ([]dao.Words, error) {
	return w.dao.GetWordsList(page, pageSize)
}

// CreateWord create word
func (w *WordsService) CreateWord(word dao.Words) error {
	res,err := w.dao.CreateWord(word)
	if err != nil {
		return err
	}
	if res > 0 {
		return nil
	}
	return errors.New("数据创建失败")
}

// UpdateWord update word by words_id
func (w *WordsService) UpdateWord(word dao.Words) error {
	res,err := w.dao.UpdateWord(word)
	if err != nil {
		return err
	}
	if res > 0 {
		return nil
	}
	return errors.New("数据更新失败，相应行数为 0")
}