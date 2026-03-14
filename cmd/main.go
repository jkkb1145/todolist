package main

import (
	"todo_list/ListenAndServer"
	"todo_list/conf"
	"todo_list/database"
)

func main() {
	conf.InitConf()
	database.ConnetDb()
	database.ConnectRedis()
	ListenAndServer.OpenSever()
}
