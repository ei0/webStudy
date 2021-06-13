package handle

import (
	"fmt"
	"math/rand"
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
	var url = RandStringBytesMaskImpr(18)
	if err != nil {
		fmt.Println("没有文件！这条可能是文本信息")
		var menu = sql.Menu{
			Uid:      id,
			Describe: c.PostForm("textarea"),
			Kind:     kind,
			URL:      url,
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
				"url":    url,
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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImpr(n int) string {
	rand.Seed(time.Now().Unix())
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return "/menu/" + string(b)
}
