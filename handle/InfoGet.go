package handle

import (
	"sql"

	"github.com/gin-gonic/gin"
)

func InfoGet(c *gin.Context) {
	id := c.Query("id")
	var user sql.User

	// intID, _ := strconv.Atoi(id)
	// menu.ID = uint(intID)
	// result := sql.SDB.Debug().Delete(&menu)
	result := sql.SDB.Where("id=?", id).First(&user)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
	} else {
		c.JSON(200, gin.H{
			"user":   user,
			"status": true,
		})
	}
}
