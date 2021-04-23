package handle

import (
	"fmt"
	"sql"

	"github.com/gin-gonic/gin"
)

func RegisterPost(c *gin.Context) {
	var user sql.User
	c.ShouldBind(&user)
	var a = sql.User{
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
	}
	fmt.Println("-------------------")
	fmt.Println(user)
	fmt.Println("-------------------")
	fmt.Println(a)
	result := sql.SDB.Debug().Create(&user)
	//fmt.Println(result)
	if result.Error != nil {
		fmt.Println("进入到注册失败返回")
		c.JSON(200, gin.H{
			"status": false,
		})
	} else {
		fmt.Println("进入到注册成功返回")
		c.JSON(200, gin.H{
			"status": true,
		})
	}

}
