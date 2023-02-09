package redisUtil

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// 声明一个全局的rdb变量
//var rdb *redis.Client
var ctx = context.Background()

// 初始化连接
func InitClient() (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
		return err
	}
	return nil
}

//func SetRbdData(key, value string) {
//	err := InitClient()
//	if err != nil {
//		return
//	}
//	//Set方法的最后一个参数表示过期时间，0表示永不过期
//	err = rdb.Set(ctx, key, value, 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	//key2将会在两分钟后过期失效
//	//err = rdb.Set( key, value, time.Minute*2).Err()
//	//if err != nil {
//	//	panic(err)
//	//}
//}
//func GetRbdData(key string) (oldCode string) {
//	err := InitClient()
//	if err != nil {
//		return
//	}
//	val, err := rdb.Get(ctx, key).Result()
//	if err == redis.Nil {
//		fmt.Println("key不存在")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Printf("值为: %v\n", val)
//	}
//	return val
//}
