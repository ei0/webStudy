package handle

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ImagePost(c *gin.Context) {
	fmt.Println("我进来了")
	file, err := c.FormFile("file")

	fmt.Println(c.PostForm("id"))
	fmt.Println(c.PostForm("textarea"))
	if err != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
		fmt.Println(err)
		return
	} else {
		fmt.Println("没报错")
	}

	//fmt.Println(file.Filename)
	dst := fmt.Sprintf("D:/Golang_test/src/01-init/static/%s", file.Filename)
	c.SaveUploadedFile(file, dst)
	// form, _ := c.MultipartForm()
	// files := form.File["upload[]"]

	// for _, file := range files {
	// 	log.Println(file.Filename)
	// 	dst := fmt.Sprintf("D:/Golang_test/src/tmp/%s", file.Filename)
	// 	//上传文件到指定的路径
	// 	c.SaveUploadedFile(file, dst)
	// }
	//c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	c.JSON(200, gin.H{
		"status": true,
	})
}
