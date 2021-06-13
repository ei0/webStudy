package handle

import (
	"fmt"
	"sql"

	"github.com/gin-gonic/gin"
)

func DetailsGet(c *gin.Context) {
	id := c.Query("id")
	var menu sql.Menu
	fmt.Println("sssss", id)
	//intID, _ := strconv.Atoi(id)
	//message.ID = uint(intID)
	result := sql.SDB.Debug().Where("id=?", id).First(&menu)
	if result.Error != nil {
		fmt.Println("失败！")
		c.JSON(200, gin.H{
			"status": false,
		})
	} else {
		fmt.Println("success！")
		var image sql.ImageURL
		sql.SDB.Where("did=?", id).Find(&image)
		imageURL := image.Url
		var user sql.User
		sql.SDB.Where("id=?", menu.Uid).First(&user)
		name := user.Name //绑定nickname
		c.JSON(200, gin.H{
			"menu":   menu,
			"status": true,
			"image":  imageURL,
			"name":   name,
		})
	}
}
