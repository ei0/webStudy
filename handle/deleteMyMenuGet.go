package handle

import (
	"sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteMyMenuGet(c *gin.Context) {
	id := c.Query("id")
	var menu sql.Menu

	intID, _ := strconv.Atoi(id)
	menu.ID = uint(intID)
	result := sql.SDB.Debug().Delete(&menu)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
	} else {
		c.JSON(200, gin.H{
			"status": true,
		})
	}

}
