package handle

import "github.com/gin-gonic/gin"

func InitMesgGet(c *gin.Context) {
	msgurl := [...]string{"http1", "http2"}
	c.JSON(200, gin.H{
		"status":  true,
		"sumMesg": 999,
		"imgurl":  msgurl,
	})

}
