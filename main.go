package main

import (
	"fmt"
	"net/http"

	"handle"
	"middleware"
	"sql"
	"tool"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// type User struct {
// 	Name      string `json:"name" form:"name"`
// 	Username  string `json:"username" form:"username"`
// 	Password  string `json:"password" form:"password"`
// 	Age       int
// 	Signature string
// 	Sex       int
// }

// func loginPost(c *gin.Context) {

// 	var user sql.User

// 	c.ShouldBind(&user)

// 	fmt.Println(user)
// 	//拿到Username和password
// 	//前往数据库验证
// 	//验证成功返回200，true
// 	//验证失败返回200, false
// 	res := db.Debug().Where("username = ? AND password = ?", user.Username, user.Password).Find(&user).RecordNotFound()
// 	fmt.Println(user)

// 	if !res {
// 		c.JSON(200, gin.H{
// 			"status":    true,
// 			"age":       user.Age,
// 			"sex":       user.Sex,
// 			"name":      user.Name,
// 			"signature": user.Signature,
// 		})
// 	} else {
// 		c.JSON(200, gin.H{
// 			"status": false,
// 		})
// 	}
// }
// func registerPost(c *gin.Context) {
// 	var user sql.User
// 	c.ShouldBind(&user)
// 	var a = sql.User{
// 		Name:     user.Name,
// 		Username: user.Username,
// 		Password: user.Password,
// 	}
// 	fmt.Println("-------------------")
// 	fmt.Println(user)
// 	fmt.Println("-------------------")
// 	fmt.Println(a)
// 	result := db.Debug().Create(&user)
// 	//fmt.Println(result)
// 	if result.Error != nil {
// 		fmt.Println("进入到注册失败返回")
// 		c.JSON(200, gin.H{
// 			"status": false,
// 		})
// 	} else {
// 		fmt.Println("进入到注册成功返回")
// 		c.JSON(200, gin.H{
// 			"status": true,
// 		})
// 	}

// }

func main() {

	db = sql.SDB

	defer db.Close()

	fmt.Println("aaa")
	fw := gin.Default()
	fw.Use(middleware.Cors())
	//fw.MaxMultipartMemory = 8 << 20
	fw.StaticFS("/static", http.Dir("./static"))
	// fw.GET("/a1", func(context *gin.Context) {
	// 	//context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	fmt.Println("请求路径：", context.FullPath())
	// 	nae := context.DefaultQuery("name", "hell")
	// 	fmt.Println("name参数: ", name)
	// 	lgendData := []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
	// 	xAxisData := []int{120, 24, rand.Intn(500), rand.Intn(500), 150, 230, 180}
	// 	context.JSON(200, gin.H{
	// 		"legend_data": legendData,
	// 		"xAxis_data":  xAxisData,
	// 	})
	// 	context.Writer.Write([]byte("Hello hi!" + name + "\n"))
	// })
	fw.POST("/login", handle.LoginPost)
	//fw.POST("/login", loginPost)

	fw.POST("/register", handle.RegisterPost)

	fw.POST("/alterinfo", handle.AlterInfoPost)
	fw.POST("/message", handle.MessagePost)
	fw.GET("/initmesg", handle.InitMesgGet)
	fw.GET("/initcharts", handle.InitChartsGet)
	fw.Run(tool.FoodWebCfg.Host + ":" + tool.FoodWebCfg.Port)
}
