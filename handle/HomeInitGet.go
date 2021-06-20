package handle

import (
	"fmt"
	"sql"

	"github.com/gin-gonic/gin"
)

func HomeInitGet(c *gin.Context) {
	var likeT []sql.Message
	var scoreT []sql.Menu
	var menusImages []string
	var nicknames []string
	//sql.SDB.Table("messages").Order("`like`").Limit(10).Find(&likeT)
	result := sql.SDB.Raw("SELECT * FROM test.messages order by `like` desc").Limit(4).Scan(&likeT)
	if result.Error != nil {
		fmt.Println("失败！", result.Error)
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	result = sql.SDB.Raw("SELECT * FROM test.menus order by `score` desc").Limit(4).Scan(&scoreT)
	if result.Error != nil {
		fmt.Println("失败！", result.Error)
		c.JSON(200, gin.H{
			"status": false,
		})
	} else {
		fmt.Println("success！")
		for _, v := range scoreT {
			var image sql.ImageURL
			sql.SDB.Where("did=?", v.ID).First(&image)
			menusImages = append(menusImages, image.Url)
		}
		for _, v := range likeT {

			var user sql.User
			sql.SDB.Where("id=?", v.Uid).First(&user)
			if user.Name == "" {
				nicknames = append(nicknames, "zhao")
			} else {
				// nicknamesLike[k] = user.Name //绑定nickname
				nicknames = append(nicknames, user.Name)
			}
		}
		c.JSON(200, gin.H{
			"menus":       scoreT,
			"menusImages": menusImages,
			"messages":    likeT,
			"status":      true,
			"nicknames":   nicknames,
		})
	}
}
