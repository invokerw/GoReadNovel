package redis

import (
	"gopkg.in/redis.v4"
)

var client *redis.Client

// 创建 redis 客户端
func createClient() *redis.Client {
	logger.ALogger().Debug("Create Redis Client..")
	client := redis.NewClient(&redis.Options{
		Addr:     "fsnsaber.cn:6379",
		Password: "weifei",
		DB:       0,
		PoolSize: 20, //连接池大小，默认为10
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	checkErr(err)
	return client
}

func init() { //如果改成int()会自动运行
	createClient()
}
func GetRedisClient() *redis.Client {
	if client == nil {
		createClient()
	}
	return client
}
func checkErr(err error) bool {
	if err != nil {
		logger.ALogger().Error(err)
		return false
	}
	return true
}
