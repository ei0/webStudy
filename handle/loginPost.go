package handle

import (
	"fmt"
	"sql"

	"github.com/gin-gonic/gin"
)

func LoginPost(c *gin.Context) {

	var user sql.User

	c.ShouldBind(&user)

	fmt.Println(user)
	//拿到Username和password
	//前往数据库验证
	//验证成功返回200，true
	//验证失败返回200, false
	res := sql.SDB.Debug().Where("username = ? AND password = ?", user.Username, user.Password).Find(&user).RecordNotFound()
	fmt.Println(user)

	if !res {
		c.JSON(200, gin.H{
			"status":    true,
			"age":       user.Age,
			"sex":       user.Sex,
			"name":      user.Name,
			"signature": user.Signature,
			"id":        user.ID,
		})
	} else {
		c.JSON(200, gin.H{
			"status": false,
		})
	}
}
