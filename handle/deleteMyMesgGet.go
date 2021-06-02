package handle

import (
	"sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteMyMesgGet(c *gin.Context) {
	id := c.Query("id")
	var message sql.Message
	intID, _ := strconv.Atoi(id)
	message.ID = uint(intID)
	result := sql.SDB.Debug().Delete(&message)
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
