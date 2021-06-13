package handle

import (
	"sql"

	"github.com/gin-gonic/gin"
)

type router struct {
	Path      string
	Component string
}

func RoutersGet(c *gin.Context) {
	// var routers []router
	// var menus []sql.Menu
	id := c.Query("id")
	var menu sql.Menu
	// //intID, _ := strconv.Atoi(id)
	// //message.ID = uint(intID)
	result := sql.SDB.Debug().Where("id=?", id).Find(&menu)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
	} else {
		c.JSON(200, gin.H{
			"status": true,
			"menu":   menu,
		})
	}

}
