package sql

import (
	"fmt"

	"tool"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name" form:"name"`
	Username  string `gorm:"type:varchar(20);unique;not null"json:"username" form:"username"`
	Password  string `gorm:"type:varchar(20);unique;not null"json:"password" form:"password"`
	Sex       string `gorm:"type:BOOLEAN; DEFAULT:true"json:"sex" form:"sex"`
	Age       string `gorm:"type:smallint; DEFAULT:18"json:"age" form:"age"`
	Signature string `gorm:"type:TEXT"json:"signature" form:"signature"`
	Id        string `json:"id" form:"id"`
}

// type Counter
var SDB *gorm.DB

func init() {
	var err error
	SDB, err = gorm.Open("mysql", tool.FoodWebCfg.MySQLLink)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("数据库连接成功")
	SDB.Debug().AutoMigrate(&User{})
	// var a = User{Name: "ni", Username: "zzk123", Password: "123"}
	// SDB.Debug().Create(&a)
}
