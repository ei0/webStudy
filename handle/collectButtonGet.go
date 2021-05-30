package handle

import (
	"fmt"
	"sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CollectButtonGet(c *gin.Context) {
	uid := c.Query("Uid")

	ty, err1 := strconv.Atoi(c.Query("type"))
	if err1 != nil {
		fmt.Println(err1)
	}
	var collect sql.Collect
	add := true
	fmt.Println("进来了，我的Uid和类型分别是", uid, ty)
	if ty == 0 {
		mid := c.Query("Mid")

		fmt.Println("进入到typo为0，我的mid是", mid)
		collect.Uid = uid
		collect.Mid = mid
		result := sql.SDB.Debug().Where("uid=? AND mid=?", uid, mid).First(&collect)
		if result.RowsAffected == 0 {
			fmt.Println("本次查询结果为0，可以进行收藏")
			sql.SDB.Debug().Create(&collect)
			var message sql.Message
			sql.SDB.Debug().Where("id=?", mid).First(&message)
			message.Collect++
			sql.SDB.Debug().Save(&message)
			add = true
		} else {
			fmt.Println("本次查询结果为1，不可以进行收藏")
			sql.SDB.Debug().Delete(&collect)
			var message sql.Message
			sql.SDB.Debug().Where("id=?", mid).First(&message)
			message.Collect--
			sql.SDB.Debug().Save(&message)
			add = false
		}
	} else if ty == 1 {
		did := c.Query("Did")
		fmt.Println("进入到typo为1，我的did是", did)
		collect.Uid = uid
		collect.Did = did
		result := sql.SDB.Debug().Where("uid=? AND did=?", uid, did).First(&collect)
		if result.RowsAffected == 0 {
			fmt.Println("本次查询结果为0，可以进行收藏")
			sql.SDB.Debug().Create(&collect)
			add = true
		} else {
			fmt.Println("本次查询结果为1，不可以进行收藏")
			sql.SDB.Debug().Delete(&collect)
			add = false
		}
	} else {
		//sid := c.Query("Sid")
	}

	c.JSON(200, gin.H{
		"status": true,
		"add":    add,
	})
}
