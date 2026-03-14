package database

import (
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"todo_list/conf"
	"todo_list/global"
)

func ConnetDb() {
	info := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Address, conf.MySqlPort, conf.DbName)
	db, err := sql.Open("mysql", info)
	if err != nil {
		panic(err)
	}
	global.Db = db
	//testDbConnect()
}

func testDbConnect() {
	if err := global.Db.Ping(); err != nil {
		fmt.Println("Database connection error")
		os.Exit(-1)
	}
	fmt.Println("Database connection successful")
}

func ConnectRedis() {
	addr := fmt.Sprintf("%s:%s", conf.RAddress, conf.RPort)
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.RPassword,
		DB:       conf.RDb,
	})
}
