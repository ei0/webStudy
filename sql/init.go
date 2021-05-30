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
type Message struct {
	gorm.Model
	Uid     string `json:"uid" form:"uid" gorm:"type:int;not null"`
	Text    string `json:"text" form:"text" gorm:"type:varchar(200)"`
	Like    int    `gorm:"type:int;DEFAULT:0"`
	Collect int    `gorm:"type:int;DEFAULT:0"`
}
type ImageURL struct {
	gorm.Model
	Mid string `json:"mid" form:"mid" gorm:"type:int;not null;DEFAULT:0"`
	Sid string `json:"sid" form:"sid" gorm:"type:int;not null;DEFAULT:0"`
	Did string `json:"did" form:"did" gorm:"type:int;not null;DEFAULT:0"`
	Url string `json:"url" form:"url" gorm:"type:varchar(255)"`
}
type Menu struct {
	gorm.Model
	Uid      string `json:"uid" form:"uid" gorm:"type:int;not null"`
	Describe string `gorm:"type:TEXT"json:"describe" form:"describe"`
	Score    int    `json:"score" form:"score" gorm:"type:int;not null;DEFAULT:0"`
	Kind     string `json:"kind" form:"kind" gorm:"type:varchar(20);not null;DEFAULT:0"`
}
type Shop struct {
	gorm.Model
	Uid      string `json:"uid" form:"uid" gorm:"type:int;not null"`
	Describe string `gorm:"type:TEXT"json:"describe" form:"describe"`
	Score    int    `json:"score" form:"score" gorm:"type:int;not null;DEFAULT:0"`
}
type Like struct {
	gorm.Model
	Mid string `json:"mid" form:"mid" gorm:"type:int;not null;DEFAULT:0"`
	Sid string `json:"sid" form:"sid" gorm:"type:int;not null;DEFAULT:0"`
	Did string `json:"did" form:"did" gorm:"type:int;not null;DEFAULT:0"`
	Uid string `json:"uid" form:"uid" gorm:"type:int;not null;DEFAULT:0"`
}

type Collect struct {
	gorm.Model
	Mid string `json:"mid" form:"mid" gorm:"type:int;not null;DEFAULT:0"`
	Sid string `json:"sid" form:"sid" gorm:"type:int;not null;DEFAULT:0"`
	Did string `json:"did" form:"did" gorm:"type:int;not null;DEFAULT:0"`
	Uid string `json:"uid" form:"uid" gorm:"type:int;not null;DEFAULT:0"`
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
	SDB.Debug().AutoMigrate(&Message{})
	SDB.Debug().AutoMigrate(&ImageURL{})
	SDB.Debug().AutoMigrate(&Menu{})
	SDB.Debug().AutoMigrate(&Collect{})
	SDB.Debug().AutoMigrate(&Like{})
	SDB.Debug().AutoMigrate(&Shop{})
	// var a = User{Name: "ni", Username: "zzk123", Password: "123"}
	// SDB.Debug().Create(&a)
}
