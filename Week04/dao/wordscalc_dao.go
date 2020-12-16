package dao

import (
	"fmt"
	"ppwords/models"
)

type CalcDb struct {
	DbType int
	Cnt int
}

var (
	w            Words
	wordsCalc    models.WordsCalc
	calculations []models.WordsCalc
	wordsList    []Words
	wordsIds     []int
)

var count int = 30

func GetWordsByUser(userId int) []Words {
	if userId == 0 {
		return []Words{}
	}
	wordsDbList := getMyDb(userId)
	if wordsDbList[2] >= 10  { // 测试库临界值 10
		if count- len(wordsList) >= 30 {
			return wordsList
		}
	} else if wordsDbList[3] >= 30 { // 复习库临界值 30
		if count- len(wordsList) >= 30 {
			return wordsList
		}
	} else if wordsDbList[1] > 0 { // 常规库不为空
		if count- len(wordsList) >= 30 {
			return wordsList
		}
	}

	// 取新词
	// TODO
	wordsIds = getMyWordsIds(userId)
	wordsList = append(wordsList, w.GetListExceptIds(wordsIds, count)...)
	if count- len(wordsList) >= 30 {
		return wordsList
	}

	// 新词也没有
	if wordsDbList[3] > 0 { // 取复习
		if count- len(wordsList) >= 30 {
			return wordsList
		}
	}
	if wordsDbList[1] > 0 { // 取常规
		if count- len(wordsList) >= 30 {
			return wordsList
		}
	}
	if wordsDbList[2] > 0 { // 取测试
		if count- len(wordsList) >= 30 {
			return wordsList
		}
	}
	if wordsDbList[4] > 0 { // 取深度1
		if count- len(wordsList) >= 30 {
			return wordsList
		}
	}
	if wordsDbList[5] > 0 { // 取深度2
		if count- len(wordsList) >= 30 {
			return wordsList
		}
	}
	if wordsDbList[6] > 0 { // 取深度3
		if count- len(wordsList) >= 30 {
			return wordsList
		}
	}

	rows,err := DB.Query("SELECT `word_id` FROM `words_calc` WHERE `user_id` = ?", userId)
	if err != nil {
		fmt.Println(err)
		return []Words{}
	}
	for rows.Next() {
		err := rows.Scan(&wordsCalc.WordId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		wordsIds = append(wordsIds, wordsCalc.WordId)
	}

	wordsList = w.GetRand(wordsIds)
	cnt := 30 - len(wordsList)
	if cnt > 0 {

	}
	return wordsList
}

func UpdateWordsByUser(data []models.WordsCalc, userId int) {
	if len(data) == 0 || userId == 0 {
		return
	}
	var (
		wordsIds []interface{}
		wordsIdsUsed []int
	)
	for _,item := range data {
		wordsIds = append(wordsIds, item.WordId)
	}
	query := fmt.Sprintf("SELECT `id`, `word_id` FROM `words_calc` WHERE word_id IN (%s) AND `user_id` = ?", Placeholders(len(wordsIds)))
	params := wordsIds
	params = append(params, userId)
	rows,err := DB.Query(query, params...)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&wordsCalc.Id, &wordsCalc.WordId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		wordsIds = append(wordsIds, wordsCalc.WordId)
	}

	fmt.Println(wordsIdsUsed, wordsIds)
	for _,row := range data {
		row.UserId = userId
		row.Click = 1

		if len(calculations) == 0 {
			DB.Exec("INSERT `words_calc` (`word_id`, `user_id`, `click`) value (?, ?, ?)", &row.WordId, &row.UserId, &row.Click)
		} else {
			for _,item := range calculations {
				fmt.Println(row)
				if item.WordId == row.WordId {
					wordsIdsUsed = append(wordsIdsUsed, item.WordId)
					row.Click = item.Click + 1
					DB.Exec("UPDATE `words_calc` SET `click` = ? WHERE `id` = ?", &row.Click, &row.Id)
				} else {
					DB.Exec("INSERT `words_calc` (`word_id`, `user_id`, `click`) value (?, ?, ?)", &row.WordId, &row.UserId, &row.Click)
				}
			}
		}
	}
}

/**
 * 获取当前用户词库数量
 */
func getMyDb(userId int) [7]int {
	rows,err := DB.Query("SELECT `db_type`, count(*) AS cnt FROM `words_calc` WHERE `user_id` = ? GROUP BY `db_type`", userId)
	if err != nil {
		panic(err)
	}
	var (
		wordsDb     CalcDb
		wordsDbList [7]int
	)
	for rows.Next() {
		err := rows.Scan(&wordsDb.DbType, &wordsDb.Cnt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		wordsDbList[wordsDb.DbType] = wordsDb.Cnt
	}
	return wordsDbList
}

/**
 * 获取当前用户已经有的词 ID
 */
func getMyWordsIds(userId int) []int {
	rows,err := DB.Query("SELECT `words_id` FROM `words_calc` WHERE `user_id` = ? ", userId)
	if err != nil {
		panic(err)
	}
	var wordsIds []int
	for rows.Next() {
		err := rows.Scan(&wordsCalc.WordId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		wordsIds = append(wordsIds, wordsCalc.WordId)
	}
	return wordsIds
}