package handle

import (
	"fmt"
	"sql"

	"github.com/gin-gonic/gin"
)

func InitChartsGet(c *gin.Context) {

	//去数据库中找到点赞数前十，收藏数前十的动态数据
	//点赞数，用户名，动态简单信息
	const LIMIT int = 10
	var likes [LIMIT]int
	var nicknames [LIMIT]string
	var messages [LIMIT]string
	var likeT []sql.Message
	//sql.SDB.Table("messages").Order("`like`").Limit(10).Find(&likeT)
	sql.SDB.Debug().Raw("SELECT * FROM test.messages order by `like` desc").Limit(10).Scan(&likeT)
	for k, v := range likeT {
		likes[k] = v.Like
		//fmt.Println(SubstrByByte(v.Text, 50))
		src := []rune(v.Text)
		if len(src) < 10 {
			messages[k] = string(src[:len(src)])
			//SubstrByByte(v.Like,10)
		} else {
			messages[k] = string(src[:10])
		}
		var user sql.User
		sql.SDB.Where("id=?", v.Uid).First(&user)
		nicknames[k] = user.Name

	}
	fmt.Println("点赞数组", likes)
	fmt.Println("用户名数组", nicknames)
	fmt.Println("消息数组", messages)
	// c.JSON(200, gin.H{
	// 	"status": true,
	// })
	c.JSON(200, gin.H{
		"status":    true,
		"likes":     likes,
		"nicknames": nicknames,
		"messages":  messages,
	})

}

//从中英混编字符串中截取length长度的字符串，不乱码，length根据byte计算
func SubstrByByte(str string, length int) string {
	if length > len([]byte(str)) {
		length = len([]byte(str))
	}
	bs := []byte(str)[:length]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}
