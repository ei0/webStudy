package handle

import (
	"fmt"
	"sql"

	"github.com/gin-gonic/gin"
)

func MyLikeGet(c *gin.Context) {
	var messages []sql.Message
	var menus []sql.Menu
	var likes []sql.Like
	uid := c.Query("uid")
	result := sql.SDB.Debug().Where("uid=?", uid).Find(&likes)
	if result.Error != nil {
		fmt.Println("失败！")
		c.JSON(200, gin.H{
			"status": false,
		})
	} else {
		fmt.Println("success！")
		for _, v := range likes {
			if v.Did == "0" {
				//这是mid 消息点赞
				var message sql.Message
				sql.SDB.Where("id=?", v.Mid).First(&message)
				messages = append(messages, message)
			} else {
				//这是did 菜单点赞
				var menu sql.Menu
				sql.SDB.Where("id=?", v.Did).First(&menu)
				menus = append(menus, menu)
			}
		}
		c.JSON(200, gin.H{
			"messages": messages,
			"status":   true,
			"menus":    menus,
		})
	}
}
