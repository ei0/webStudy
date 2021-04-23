package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FoodWebConfig struct {

	//软件名
	Name string
	//gin模式
	Mode string
	//主机IP
	Host string
	//主机端口
	Port string
	//配置文件路径
	ConFilePath string
	//数据库连接
	MySQLLink string
}

//全局变量实例
var FoodWebCfg *FoodWebConfig

//判断一个文件是否存在
func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)

	if err == nil {

		return true, nil
	}
	if os.IsNotExist(err) {
		fmt.Println("服务器打开路径失败")
		return false, nil
	}
	return false, err
}

//读取用户配置文件

func (g *FoodWebConfig) Reload() {

	if confFileExists, _ := PathExists(g.ConFilePath); confFileExists != true {

		return
	}
	//从路径中读取文件
	data, err := ioutil.ReadFile(FoodWebCfg.ConFilePath)
	if err != nil {
		panic(err)
	}

	//把json数据解析到struct中
	err = json.Unmarshal(data, &FoodWebCfg)
	if err != nil {
		panic(err)
	}

}

func init() {

	//初始化GlobalObject变量
	FoodWebCfg = &FoodWebConfig{
		ConFilePath: "./config/foodWeb.json",
	}

	//从文件中读取路径并且把值读入
	FoodWebCfg.Reload()
}
