package main

import (
	"Notes-go-project/api"
	Mg "Notes-go-project/manage_socket_conn"
	"Notes-go-project/utility/databaseConnection"
)

func main() {
	//连接数据库
	databaseConnection.GetDB()
	//路由

	go Mg.GetCharRoomThread().Start()
	api.LaunchProject()
}
