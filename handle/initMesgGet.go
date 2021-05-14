package handle

import (
	"fmt"
	"sql"

	"github.com/gin-gonic/gin"
)

func InitMesgGetA(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
	})
}

func InitMesgGet(c *gin.Context) {
	var messages []sql.Message
	var count int
	const LIMIT int = 10
	//拿到数据库中所有的动态消息,message内容，创建时间，用户id
	sql.SDB.Table("messages").Limit(LIMIT).Find(&messages).Limit(-1).Count(&count)
	var uids [LIMIT]string
	var nicknames [LIMIT]string
	var creatTimes [LIMIT]string
	var Texts [LIMIT]string
	var imageCounts [LIMIT]int
	var imageURLs [LIMIT][]string
	var likeCounts [LIMIT]int
	var collectCounts [LIMIT]int

	for k, v := range messages {
		fmt.Println("第%d次访问messages", k)
		uids[k] = v.Uid //绑定uid
		fmt.Println("第%d次,绑定uid,原始值%s,绑定后%s", k, v.Uid, uids[k])
		fmt.Println("第")
		creatTimes[k] = v.CreatedAt.Format("2006-01-02 15:04:05") //绑定创建时间
		fmt.Println("第%d次,绑定创建时间,原始值%s,绑定后%s", k, v.CreatedAt, creatTimes[k])

		Texts[k] = v.Text //绑定消息内容
		fmt.Println("第%d次,绑定消息内容,原始值%s,绑定后%s", k, v.Text, Texts[k])

		likeCounts[k] = v.Like
		fmt.Println("第%d次,绑定消息内容,原始值%s,绑定后%s", k, v.Like, likeCounts[k])
		collectCounts[k] = v.Collect
		fmt.Println("第%d次,绑定消息内容,原始值%s,绑定后%s", k, v.Collect, collectCounts[k])

		var user sql.User
		sql.SDB.Where("id=?", v.Uid).First(&user)
		nicknames[k] = user.Name //绑定nickname
		fmt.Println("第%d次,绑定nickname,原始值%s,绑定后%s", k, user.Name, nicknames[k])

		var images []sql.ImageURL
		sql.SDB.Where("mid=?", v.ID).Find(&images).Count(&imageCounts[k])
		for _, v := range images {
			imageURLs[k] = append(imageURLs[k], v.Url) //绑定图片
		}
		fmt.Printf("第%d次,绑定图片url,绑定后", k)
		fmt.Println(imageURLs[k])
		fmt.Println("第%d次访问messages,结束————————————————————", k)
	}

	//fmt.Println(messages, count)
	//需要返回的数据，动态总数+动态{用户id，用户名，创建时间，消息内容，图片个数，图片路径[]，点赞数，收藏数}
	//msgurl := [...]string{"http1", "http2"}

	c.JSON(200, gin.H{
		"status":        true,
		"messageCount":  count,
		"uids":          uids,
		"nicknames":     nicknames,
		"creatTimes":    creatTimes,
		"Texts":         Texts,
		"imageCounts":   imageCounts,
		"imageURLs":     imageURLs,
		"likeCounts":    likeCounts,
		"collectCounts": collectCounts,
	})
	//返回动态总数目

}
