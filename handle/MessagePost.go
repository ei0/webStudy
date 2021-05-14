package handle

import (
	"fmt"
	"sql"
	"time"

	"github.com/gin-gonic/gin"
)

func MessagePost(c *gin.Context) {
	fmt.Println("我进来了")
	//form, err := c.MultipartForm()
	id := c.PostForm("id")
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("没有文件！这条可能是文本信息")
		var mesg = sql.Message{
			Uid:  id,
			Text: c.PostForm("textarea"),
		}
		result := sql.SDB.Debug().Create(&mesg)
		if result.Error != nil {
			c.JSON(200, gin.H{
				"status": false,
			})
		} else {
			sql.SDB.Debug().Last(&mesg)
			//fmt.Println(mesg.ID)
			//fmt.Println(mesg.Uid)
			c.JSON(200, gin.H{
				"status": true,
				"mid":    mesg.ID,
			})
		}

	} else {
		fmt.Println("这条是个带文件的post")
		imageURL := fmt.Sprintf("%d%s", time.Now().UnixNano(), file.Filename)
		dst := fmt.Sprintf("D:/Golang_test/src/01-init/static/%s", imageURL)
		c.SaveUploadedFile(file, dst)
		var imgurl = sql.ImageURL{
			Mid: id,
			Url: imageURL,
		}
		result := sql.SDB.Debug().Create(&imgurl)
		if result.Error != nil {
			fmt.Println(result.Error)
			c.JSON(200, gin.H{
				"status": true,
			})

		} else {
			c.JSON(200, gin.H{
				"status": true,
			})
		}

	}
}
