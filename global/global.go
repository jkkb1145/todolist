package global

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var Db *sql.DB
var RedisClient *redis.Client

func GetDb() *sql.DB {
	return Db
}
