package handle

import (
	"fmt"
	"sql"

	"github.com/gin-gonic/gin"
)

func MyMesgsAGet(c *gin.Context) {
	var message []sql.Message
	uid := c.Query("uid")
	result := sql.SDB.Debug().Where("uid=?", uid).Find(&message)
	if result.Error != nil {
		fmt.Println("失败！")
		c.JSON(200, gin.H{
			"status": false,
		})
	} else {
		fmt.Println("success！")
		c.JSON(200, gin.H{
			"messages": message,
			"status":   true,
		})
	}
}
