package main

import (
	"go_web_demo/dao"
	"go_web_demo/routers"
)


func main() {
	//连接数据库
	err := dao.InitMysql()
	if err != nil{
		panic(err)
	}
	defer dao.Close()
	//模型绑定
	dao.InitModel()
	r := routers.SetupRouter()
	r.Run() //启动，默认监听8080端口
}
