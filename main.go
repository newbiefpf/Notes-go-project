package main

import (
	"Notes-go-project/api"
	"Notes-go-project/utility/databaseConnection"
)

func main() {
	//连接数据库
	databaseConnection.GetDB()
	//路由
	api.LaunchProject()
}
