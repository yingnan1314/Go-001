package dao

import (
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
	rand2 "math/rand"
	"ppwords/models"
)

type Words models.Words

var (
	wordsMaxId int = 5241
	words      []Words
	word       Words
	params     []interface{}
)

var backendDao interface{
	GetWordsList(page, pageSize int) []Words
	CreateWord(word Words) error
	UpdateWord(word Words) int
}

var appDao interface{
	GetRand(wordsIds []int) []Words
	GetListExceptIds(wordsIds []int, count int) []Words
}

// @Title 获取单词列表
// @description 获取单词列表
// @auth yingnan 2020/11/16 17:14
// @param page int "页码"
// @param pageSize int "每页条数"
// @return  []Words "单词列表"
func (w *Words) GetWordsList(page, pageSize int) ([]Words,error) {
	if page < 1 {
		page = 1
	}
	rows,err := DB.Query("SELECT `id`, `word`, `meaning`, `phonetic_symbol_en`, `phonetic_symbol_un`, " +
		"`partofspeech`, `grade`, `type`, `created_at`, `updated_at` FROM `words` LIMIT ? OFFSET ?", pageSize, pageSize * (page - 1))
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	wordsList := w.controlWords(rows)
	fmt.Println(wordsList)
	return wordsList,nil
}

func (w *Words) CreateWord(word Words) (int64, error) {
	res,err := DB.Exec("INSERT `words`（`word`, `meaning`, `phonetic_symbol_en`, " +
		"`phonetic_symbol_un`, `partofspeech`, `type`, `grade`）VALUE (?, ?, ?, ?, ?, ?, ?)",
		word.Word,
		word.Meaning,
		word.PhoneticSymbolEn,
		word.PhoneticSymbolUn,
		word.Partofspeech,
		word.Type,
		word.Grade)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return res.RowsAffected()
}
func (w *Words) UpdateWord(word Words) (int64, error) {
	res,err := DB.Exec("UPDATE `words` SET `word` = ?, `meaning` = ?, `phonetic_symbol_en` = ?, " +
		"`phonetic_symbol_un` = ?, `partofspeech` = ?, `type` = ?, `grade` = ? WHERE `id` = ?",
		word.Word,
		word.Meaning,
		word.PhoneticSymbolEn,
		word.PhoneticSymbolUn,
		word.Partofspeech,
		word.Type,
		word.Grade,
		word.Id)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return res.RowsAffected()
}

// @Title 随机获取单词
// @description 随机获取单词
// @auth yingnan 2020/11/16 17:14
// @param wordsIds []int "单词ID列表"
// @return  []Words "单词列表"
func (w *Words) GetRand(wordsIds []int) []Words {
	query := fmt.Sprintf("SELECT id,word,meaning,grade,type FROM words " +
		"WHERE id IN (%s) LIMIT 20", Placeholders(len(wordsIds)))
	fmt.Println(query,wordsIds)
	for w := range wordsIds {
		params = append(params, interface{}(w))
	}
	rows, err  := DB.Query(query, params...)
	if err != nil {
		panic(err)
	}
	wordsList := w.controlWords(rows)
	fmt.Println(wordsList)
	return wordsList
}

// @Title 随机获取除指定ID以外的指定数量的单词列表
// @description 随机获取除指定ID以外的指定数量的单词列表
// @auth yingnan 2020/11/16 17:14
// @param wordsIds []int "单词ID列表"
// @param count int "指定数量"
// @return  []Words "单词列表"
func (w *Words) GetListExceptIds(wordsIds []int, count int) []Words {
	var (
		limit int
		flag bool
	)
	if len(wordsIds) + count > wordsMaxId {
		limit = 0
		flag = false
	} else {
		limit = rand2.Intn(wordsMaxId)
		flag = true
	}

	query := fmt.Sprintf("SELECT id,word,meaning,grade,type FROM words " +
		"WHERE id IN (%s) LIMIT %d OFFSET %d", Placeholders(len(wordsIds)), limit, count)
	for w := range wordsIds {
		params = append(params, interface{}(w))
	}
	params = append(params, limit, count)
	rows,err := DB.Query(query, params)
	if err != nil {
		panic(err)
	}
	wordsList := w.controlWords(rows)
	if flag && len(wordsList) < count {
		w.GetListExceptIds(wordsIds, count - len(wordsList))
	}
	return wordsList
}

// @Title 整理单词列表
// @description 整理单词列表
// @auth yingnan 2020/11/16 17:14
// @param rows *sql.Rows ""
// @return  []Words "单词列表"
func (w *Words) controlWords(rows *sql.Rows) []Words {
	wordsList := make([]Words, 0)
	for rows.Next() {
		err := rows.Scan(&word.Id,
			&word.Word,
			&word.Meaning,
			&word.PhoneticSymbolEn,
			&word.PhoneticSymbolUn,
			&word.Partofspeech,
			&word.Grade,
			&word.Type,
			&word.CreatedAt,
			&word.UpdatedAt)
		if err == nil {
			wordsList = append(wordsList, word)
		}
	}
	return wordsList
}
