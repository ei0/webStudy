package handle

import (
	"fmt"
	"sql"
	"time"

	"github.com/gin-gonic/gin"
)

func MenuPost(c *gin.Context) {
	fmt.Println("我进来了")
	//form, err := c.MultipartForm()
	id := c.PostForm("id")
	file, err := c.FormFile("file")
	kind := c.PostForm("kind")
	if err != nil {
		fmt.Println("没有文件！这条可能是文本信息")
		var menu = sql.Menu{
			Uid:      id,
			Describe: c.PostForm("textarea"),
			Kind:     kind,
		}
		result := sql.SDB.Debug().Create(&menu)
		if result.Error != nil {
			c.JSON(200, gin.H{
				"status": false,
			})
		} else {
			sql.SDB.Debug().Last(&menu)
			//fmt.Println(mesg.ID)
			//fmt.Println(mesg.Uid)
			c.JSON(200, gin.H{
				"status": true,
				"mid":    menu.ID,
			})
		}

	} else {
		fmt.Println("这条是个带文件的post")
		imageURL := fmt.Sprintf("%d%s", time.Now().UnixNano(), file.Filename)
		dst := fmt.Sprintf("D:/Golang_test/src/01-init/static/%s", imageURL)
		c.SaveUploadedFile(file, dst)
		var imgurl = sql.ImageURL{
			Did: id,
			Url: imageURL,
		}
		result := sql.SDB.Debug().Create(&imgurl)
		if result.Error != nil {
			fmt.Println(result.Error)
			c.JSON(200, gin.H{
				"status": false,
			})

		} else {
			c.JSON(200, gin.H{
				"status": true,
			})
		}

	}
}
