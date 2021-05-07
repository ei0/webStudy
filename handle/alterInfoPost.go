package handle

import (
	"fmt"
	"sql"
	"tool"

	"github.com/gin-gonic/gin"
)

func AlterInfoPost(c *gin.Context) {
	var user sql.User
	c.ShouldBind(&user)
	fmt.Println(user)
	tool.PrintStruct(user)
	result := sql.SDB.Debug().Model(&user).Where("id = ?", user.Id).Updates(sql.User{Name: user.Name, Age: user.Age, Sex: user.Sex, Signature: user.Signature})

	if result.Error != nil {
		fmt.Println("更新个人信息出错")
		c.JSON(200, gin.H{
			"status": false,
		})
	} else {
		fmt.Println("更新个人信息成功")
		c.JSON(200, gin.H{
			"status": true,
		})
	}

}
