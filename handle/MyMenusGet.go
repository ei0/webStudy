package handle

import (
	"fmt"
	"sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MyMenusGet(c *gin.Context) {

	var page = c.Query("id")
	fmt.Println("id==", page)
	var kind = c.Query("kind")
	fmt.Println("kind==", kind)
	var uid = c.Query("uid")
	var menu []sql.Menu
	var count int
	var curCount int
	const LIMIT int = 9
	var imageURLs [LIMIT]string
	var likeCount [LIMIT]int
	var collectCount [LIMIT]int
	var flageLikes [LIMIT]int
	var flageCollects [LIMIT]int
	val, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println(err)
	}
	var start = (val - 1) * LIMIT
	fmt.Println("从第几条开始查找数据库", start)
	//拿到数据库中所有的动态消息,message内容，创建时间，用户id
	sql.SDB.Debug().Table("menus").Where("kind=? AND uid=?", kind, uid).Offset(start).Limit(LIMIT).Find(&menu).Limit(-1).Count(&count)

	for k, v := range menu {

		var like sql.Like
		result := sql.SDB.Debug().Where("did=? AND uid=?", v.ID, uid).First(&like)
		if result.RowsAffected == 1 {
			flageLikes[k] = 0
		} else {
			flageLikes[k] = 1
		}
		var collect sql.Collect
		result = sql.SDB.Where("did=? AND uid=?", v.ID, uid).First(&collect)
		if result.RowsAffected == 1 {
			flageCollects[k] = 0
		} else {
			flageCollects[k] = 1
		}

		var image sql.ImageURL
		sql.SDB.Where("did=?", v.ID).Find(&image)
		imageURLs[k] = image.Url
		var likes sql.Like
		var temp int
		temp = 0
		sql.SDB.Where("did=?", v.ID).Find(&likes).Count(&temp)
		likeCount[k] = temp
		var collects sql.Collect
		sql.SDB.Where("did=?", v.ID).Find(&collects).Count(&temp)
		collectCount[k] = temp

	}
	fmt.Println("menu信息", menu)
	curCount = len(menu)
	c.JSON(200, gin.H{
		"status":       true,
		"menus":        menu,
		"counts":       count,
		"curCount":     curCount,
		"imageURLs":    imageURLs,
		"likeCount":    likeCount,
		"collectCount": collectCount,
		"flageLike":    flageLikes,
		"flageCollect": flageCollects,
	})
}
