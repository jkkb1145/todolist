package conf

import (
	"gopkg.in/ini.v1"
	"strconv"
)

var (
	Address    string
	ServerPort string
	//
	Host      string
	MySqlPort string
	User      string
	Password  string
	DbName    string
	//
	RAddress  string
	RPort     string
	RPassword string
	RDb       int
	//
	JWTKey string
)

func InitConf() {
	file, err := ini.Load("D:\\Godemo\\todo_list\\conf\\config.ini")
	if err != nil {
		panic(err)
	}
	initServer(file)
	initMySql(file)
	initRedis(file)
	initKey(file)
}

func initServer(file *ini.File) {
	Address = file.Section("server").Key("Address").String()
	ServerPort = file.Section("server").Key("Port").String()
}

func initMySql(file *ini.File) {
	Host = file.Section("mysql").Key("Host").String()
	MySqlPort = file.Section("mysql").Key("Port").String()
	User = file.Section("mysql").Key("User").String()
	Password = file.Section("mysql").Key("Password").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func initRedis(file *ini.File) {
	RAddress = file.Section("redis").Key("Address").String()
	RPort = file.Section("redis").Key("Port").String()
	RPassword = file.Section("redis").Key("Password").String()
	DbString := file.Section("redis").Key("Db").String()
	RDb, _ = strconv.Atoi(DbString)
}

func initKey(file *ini.File) {
	JWTKey = file.Section("key").Key("Jwtkey").String()
}
