package words

import (
	"github.com/gin-gonic/gin"
	"ppwords/app/service/user"
	"ppwords/dao"
	"ppwords/models"
)

func GetWordsByUser(c *gin.Context) {
	user := user.User{}.GetUserInfo()
	data := dao.GetWordsByUser(user.Id)
	c.JSON(200, data)
}

func UpdateWordsByUser(c *gin.Context) {
	dao.UpdateWordsByUser([]models.WordsCalc{{Id: 0, WordId: 1, Active: 1}}, 1)
	c.JSON(200, []string{})
}
