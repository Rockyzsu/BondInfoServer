package cache

import (
	"bondinfoserver/service"
	"fmt"
	"github.com/go-redis/redis"
)

type Cache struct {
	conn *redis.Client
}

func (this *Cache) CacheInit() {
	this.connect()

}
func (this *Cache) connect() {
	conf := service.ReadRedisConfig()
	this.conn = redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})

}
func (this *Cache) Get(id string) (string, bool) {
	result, err := this.conn.Get(id).Result()
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	return result, true
}

func (this *Cache) Set(id string, content string) bool {
	result, err := this.conn.Set(id, content, 0).Result()
	if err != nil {
		return false
	} else {
		fmt.Println("-----------", result)
		return true
	}
}

func (this *Cache) CheckUserExist(sign string) bool {
	result, err := this.conn.SRem("User", sign).Result()
	if err != nil {
		fmt.Println("获取结果失败")
		return false
	}
	if result == 0 {
		return false
	} else {
		return true
	}
}
