package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"strconv"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	//从本地读取环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadRedis(file)
	Redis()
}

func LoadRedis(file *ini.File) {
	redisSection := file.Section("redis")
	RedisDb = redisSection.Key("RedisDb").String()
	RedisAddr = redisSection.Key("RedisAddr").String()
	RedisPw = redisSection.Key("RedisPw").String()
	RedisDbName = redisSection.Key("RedisDbName").String()
}

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Error(err)
		panic(err)
	}
	RedisClient = client
}
