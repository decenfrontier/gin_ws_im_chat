package conf

import (
	"chat/model"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus" //github.com/sirupsen/logrus
	"go.mongodb.org/mongo-driver/mongo"  //MongoDB的Go驱动包
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/ini.v1"
	"strconv"
	"strings"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	MongoDBClient *mongo.Client
	MongoDBName   string
	MongoDBAddr   string
	MongoDBPwd    string
	MongoDBPort   string
)

func init() {
	//从本地读取环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)

	LoadMysqlData(file)
	LoadMongoDB(file)
	LoadRedis(file)
	// 连接MySQL
	MySQL()
	// 连接MongoDB
	MongoDB()
	// 连接Redis
	Redis()
}

func MongoDB() {
	// 设置mongoDB客户端连接信息
	clientOptions := options.Client().ApplyURI("mongodb://" + MongoDBAddr + ":" + MongoDBPort)
	var err error
	MongoDBClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logging.Info(err)
	}
	err = MongoDBClient.Ping(context.TODO(), nil)
	if err != nil {
		logging.Info(err)
	}
	logging.Info("MongoDB Connect")
}

func MySQL() {
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
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

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadMongoDB(file *ini.File) {
	MongoDBName = file.Section("MongoDB").Key("MongoDBName").String()
	MongoDBAddr = file.Section("MongoDB").Key("MongoDBAddr").String()
	MongoDBPwd = file.Section("MongoDB").Key("MongoDBPwd").String()
	MongoDBPort = file.Section("MongoDB").Key("MongoDBPort").String()
}

func LoadRedis(file *ini.File) {
	redisSection := file.Section("redis")
	RedisDb = redisSection.Key("RedisDb").String()
	RedisAddr = redisSection.Key("RedisAddr").String()
	RedisPw = redisSection.Key("RedisPw").String()
	RedisDbName = redisSection.Key("RedisDbName").String()
}
