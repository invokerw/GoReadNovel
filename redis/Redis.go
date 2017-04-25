package redis

import (
	"GoReadNovel/logger"
	"fmt"
	"gopkg.in/redis.v4"
	"os"
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

func GetGuid() string {
	f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	logger.ALogger().Debugf("Guid = ", uuid)
	return uuid
}

/*
func ExampleClient() {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exists
}
*/
